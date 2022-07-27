package mq

import (
	"os"

	"github.com/darabuchi/log"
	"github.com/darabuchi/utils"
)

type Option struct {
	// 数据存储的路径
	DataPath string `json:"data_path,omitempty"`
	// 内存队列的大小
	MemQueueSize int64 `json:"mem_queue_size,omitempty"`
}

func NewOption() *Option {
	return &Option{}
}

func (p *Option) init() error {
	if p.DataPath == "" {
		p.DataPath = utils.GetExecPath()
	}

	if !utils.FileExists(p.DataPath) {
		err := os.MkdirAll(p.DataPath, 0777)
		if err != nil {
			log.Errorf("err:%v", err)
			return err
		}
	}

	if p.MemQueueSize <= 0 {
		p.MemQueueSize = -1
	}

	return nil
}
