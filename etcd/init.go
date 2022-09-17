package etcd

import (
	"context"
	"time"

	"github.com/darabuchi/log"
	"github.com/darabuchi/utils"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
)

var cli *clientv3.Client

func Connect(c Config) error {
	var err error
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:            c.Addrs,
		AutoSyncInterval:     time.Second,
		DialTimeout:          5 * time.Second,
		DialKeepAliveTime:    time.Second,
		DialKeepAliveTimeout: time.Minute,
		MaxCallSendMsgSize:   0,
		MaxCallRecvMsgSize:   0,
		Username:             "",
		Password:             "",
		RejectOldCluster:     true,
		DialOptions:          nil,
		Context:              nil,
		Logger:               zap.NewNop(),
		PermitWithoutStream:  false,
	})
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}

	return nil
}

type EventType string

const (
	Changed EventType = "changed"
	Deleted EventType = "delete"
)

type Event struct {
	Key     string
	Type    EventType
	Value   string
	Version int64
}

func Watch(key string, logic func(event Event)) {
	wc := cli.Watch(context.Background(), key)
	go func() {
		for {
			select {
			case v := <-wc:
				for _, event := range v.Events {
					logic(Event{
						Key:   string(event.Kv.Key),
						Value: string(event.Kv.Value),

						Version: event.Kv.Version,

						Type: func() EventType {
							switch event.Type {
							case mvccpb.PUT:
								return Changed
							case mvccpb.DELETE:
								return Deleted
							default:
								return ""
							}
						}(),
					})
				}
			}
		}
	}()
}

func Set(key string, val interface{}) error {
	_, err := cli.Put(context.TODO(), key, utils.ToString(val))
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}

	return nil
}

func SetEx(key string, val interface{}, timeout time.Duration) error {
	lease := clientv3.NewLease(cli)

	leaseRsp, err := lease.Grant(context.TODO(), int64(timeout.Seconds()))
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}

	_, err = cli.Put(context.TODO(), key, utils.ToString(val), clientv3.WithLease(leaseRsp.ID))
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}

	return nil
}
