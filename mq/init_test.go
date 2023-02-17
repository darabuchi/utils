package mq

import (
	"testing"
	"time"
)

func TestMq(t *testing.T) {
	err := Start(&Option{
		DataPath:     "/temp/mq/",
		MemQueueSize: -1,
	})
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}

	type TempData struct {
		Temp int
	}

	RegisterTopic(Topic("temp"), &TempData{})
	RegisterHandel(Topic("temp"), Channel("default"), &Handle{
		HandleFunc: func(msg *HandleReq) (*HandleRsp, error) {
			var rsp HandleRsp

			req := msg.Message.(*TempData)

			t.Logf("got queue %v", req.Temp)

			return &rsp, nil
		},
	})

	for i := 0; i < 1000; i++ {
		Publish(Topic("temp"), &TempData{
			Temp: i,
		})
	}

	time.Sleep(time.Minute * 10)
}
