package mq

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/darabuchi/log"
	"github.com/darabuchi/utils"
	"github.com/nsqio/nsq/nsqd"
)

const (
	Version = 1
)

type PublishRsp struct {
	Version uint32 `json:"version,omitempty"`
	MsgId   string `json:"msgid,omitempty"`
}

func Publish(topicName fmt.Stringer, message interface{}) (*PublishRsp, error) {
	var err error

	var value []byte
	switch x := message.(type) {
	case []byte:
		value = x
	case string:
		value = []byte(x)
	default:
		err = utils.Validate(message)
		if err != nil {
			log.Errorf("err:%v", err)
			return nil, err
		}
		value, err = json.Marshal(x)
		if err != nil {
			log.Errorf("err:%v", err)
			return nil, err
		}
	}

	return pub(topicName, &nsqMsg{
		Version: Version,
		Body:    value,
	})
}

func pub(topicName fmt.Stringer, msg *nsqMsg) (*PublishRsp, error) {
	msg.PubAt = time.Now()
	msg.TraceId = log.GenTraceId()

	buf, err := json.Marshal(msg)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	topic := producer.GetTopic(topicName.String())
	mqMsg := nsqd.NewMessage(topic.GenerateID(), buf)
	log.Infof("topic:%s,id:%s,msg:%s", topicName.String(), mqMsg.ID, buf)
	err = topic.PutMessage(mqMsg)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}
	return &PublishRsp{
		Version: msg.Version,
		MsgId:   mqMsg.ID.String(),
	}, nil
}
