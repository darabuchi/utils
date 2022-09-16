package etcd

import (
	"testing"
	"time"
	
	"github.com/darabuchi/log"
)

func TestWatch(t *testing.T) {
	err := Connect(Config{
		Addrs: []string{"http://127.0.0.1:2379"},
	})
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}
	
	Watch("test", func(event Event) {
		log.Info(event.Type)
	})
	
	for i := 0; i < 10; i++ {
		Set("test", time.Now().String())
		time.Sleep(time.Second * 3)
	}
}
