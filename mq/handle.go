package mq

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sync"
	"time"

	"github.com/darabuchi/log"
	"github.com/darabuchi/utils"
	"github.com/nsqio/nsq/nsqd"
)

type HandleReq struct {
	MsgId    string    `json:"msg_id,omitempty"`
	Attempts uint16    `json:"attempts,omitempty"`
	PubAt    time.Time `json:"pub_at,omitempty"`

	Body    []byte      `json:"body,omitempty"`
	Message interface{} `json:"message,omitempty"`
}

type HandleRsp struct {
	NeedRetry bool `json:"need_retry,omitempty"`

	WaitTime time.Duration `json:"wait_time,omitempty"`
}

type HandleFunc func(msg *HandleReq) (*HandleRsp, error)

type Handle struct {
	HandleFunc HandleFunc `json:"handle_func,omitempty"`
	// 默认为无限重试
	MaxAttempts uint16 `json:"max_attempts,omitempty"`

	MaxProcessCnt int64 `json:"max_process_cnt,omitempty"`

	// 消息超时时间
	MsgTimeout time.Duration `json:"msg_timeout,omitempty"`
}

type nsqMsg struct {
	Version uint32    `json:"version,omitempty" validate:"required"`
	Body    []byte    `json:"body,omitempty" validate:"required"`
	TraceId string    `json:"trace_id,omitempty"`
	PubAt   time.Time `json:"pub_at,omitempty" validate:"required"`
}

var (
	topicType sync.Map
)

func RegisterTopic(name fmt.Stringer, messageType interface{}) {
	rt := reflect.TypeOf(messageType)

	for rt.Kind() == reflect.Ptr {
		rt = rt.Elem()
	}

	topicType.Store(name.String(), rt)
}

func RegisterHandel(topicName fmt.Stringer, channelName fmt.Stringer, handel *Handle) {
	if handel.HandleFunc == nil {
		log.Panic("HandleFunc is nil")
	}

	var rt reflect.Type
	{
		rtv, ok := topicType.Load(topicName.String())
		if !ok {
			log.Panicf("topic %s not register", topicName.String())
		}
		rt = rtv.(reflect.Type)
	}

	topic := producer.GetTopic(topicName.String())

	channel := topic.GetChannel(channelName.String())

	if handel.MaxProcessCnt < 1 {
		handel.MaxProcessCnt = 1
	}

	for i := int64(0); i < handel.MaxProcessCnt; i++ {
		go func(clientId int64) {
			logic := func(msg *nsqd.Message) {
				log.SetTrace(msg.ID.String())
				log.Infof("handel %s %s", topicName.String(), channelName.String())

				finishC := make(chan bool, 1)

				if handel.MaxAttempts > 0 && msg.Attempts > handel.MaxAttempts {
					err := channel.FinishMessage(clientId, msg.ID)
					if err != nil {
						log.Errorf("err:%v", err)
						return
					}
					return
				}

				msg.Attempts++

				err := channel.StartInFlightTimeout(msg, clientId, time.Second*12)
				if err != nil {
					log.Errorf("err:%v", err)
					return
				}

				go func() {
					ticker := time.NewTicker(time.Second * 5)
					defer ticker.Stop()
					for {
						select {
						case <-ticker.C:
							err = channel.TouchMessage(clientId, msg.ID, time.Second*12)
							if err != nil {
								log.Errorf("err:%v", err)
							}
						case <-finishC:
							return
						}
					}
				}()
				defer func() {
					finishC <- true
				}()

				defer utils.CachePanicWithHandle(func(err interface{}) {
					err = channel.FinishMessage(clientId, msg.ID)
					if err != nil {
						log.Errorf("err:%v", err)
						return
					}
				})

				var m nsqMsg
				err = json.Unmarshal(msg.Body, &m)
				if err != nil {
					log.Errorf("err:%v", err)

					err = channel.FinishMessage(clientId, msg.ID)
					if err != nil {
						log.Errorf("err:%v", err)
						return
					}

					return
				}

				if handel.MsgTimeout > 0 && time.Since(m.PubAt) > handel.MsgTimeout {
					log.Warnf("msg timeout")
					err := channel.FinishMessage(clientId, msg.ID)
					if err != nil {
						log.Errorf("err:%v", err)
						return
					}
					return
				}

				log.SetTrace(m.TraceId + "." + log.GenTraceId())
				log.Infof("[%s(%s)]handel msg %s", topicName.String(), channelName.String(), msg.ID.String())

				req := &HandleReq{
					MsgId:    msg.ID.String(),
					Attempts: msg.Attempts,
					Body:     m.Body,
					PubAt:    m.PubAt,
				}

				req.Message = reflect.New(rt).Interface()
				err = json.Unmarshal(req.Body, req.Message)
				if err != nil {
					log.Errorf("err:%v", err)
				} else {
					err = utils.Validate(req.Message)
					if err != nil {
						log.Errorf("topic:%s err:%v", topicName.String(), err)

						err = channel.FinishMessage(clientId, msg.ID)
						if err != nil {
							log.Errorf("err:%v", err)
							return
						}

						return
					}
				}

				rsp, err := handel.HandleFunc(req)
				if err != nil {
					log.Errorf("err:%v", err)

					switch x := err.(type) {
					case *utils.Error:
						if x.NeedRetry {
							// 默认重试，除非用户指定了
							if rsp == nil || rsp.NeedRetry {
								if rsp != nil && rsp.WaitTime > 0 {
									log.Warnf("retry %s after %v", msg.ID.String(), rsp.WaitTime)
									err = channel.RequeueMessage(clientId, msg.ID, rsp.WaitTime)
								} else {
									log.Warnf("retry %s after 5s", msg.ID.String())
									err = channel.RequeueMessage(clientId, msg.ID, time.Second*5)
								}
								if err != nil {
									log.Errorf("err:%v", err)
									return
								}
							}
						}
					default:
						// 除非指定重试
						if rsp != nil && rsp.NeedRetry {
							if rsp.WaitTime > 0 {
								err = channel.RequeueMessage(clientId, msg.ID, rsp.WaitTime)
							} else {
								err = channel.RequeueMessage(clientId, msg.ID, time.Second*5)
							}
							if err != nil {
								log.Errorf("err:%v", err)
								return
							}
							return
						}
					}

					return
				}

				// 需要重试
				if rsp.NeedRetry {
					if rsp.WaitTime > 0 {
						err = channel.RequeueMessage(clientId, msg.ID, rsp.WaitTime)
					} else {
						err = channel.RequeueMessage(clientId, msg.ID, time.Second*5)
					}
					if err != nil {
						log.Errorf("err:%v", err)
						return
					}
					return
				}

				// 正常结束
				err = channel.FinishMessage(clientId, msg.ID)
				if err != nil {
					log.Errorf("err:%v", err)
					return
				}
			}

			for {
				msg := channel.ReadMessage()
				if msg == nil {
					continue
				}

				logic(msg)
			}
		}(int64(i))
	}
}
