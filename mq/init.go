package mq

import (
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/darabuchi/log"
	"github.com/darabuchi/utils"
	"github.com/darabuchi/utils/xtime"
	"github.com/nsqio/nsq/nsqd"
)

var producer *nsqd.NSQD

func Start(path string) error {
	opt := nsqd.NewOptions()
	opt.HTTPAddress = ""
	opt.HTTPSAddress = ""
	opt.TCPAddress = ""
	opt.BroadcastAddress = ""
	opt.MaxMsgSize = -1
	opt.MemQueueSize = 10

	opt.DataPath = utils.GetExecPath()
	if path != "" {
		opt.DataPath = path
	}

	// opt.SnappyEnabled = true
	opt.MaxMsgSize = 1024 * 1024 * 1024

	if !utils.FileExists(opt.DataPath) {
		err := os.MkdirAll(opt.DataPath, 0777)
		if err != nil {
			log.Errorf("err:%v", err)
			return err
		}
	}

	opt.Logger = NewLogger()

	var err error
	producer, err = nsqd.New(opt)
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}

	go func() {
		err := producer.Main()
		if err != nil {
			log.Panicf("err:%v", err)
		}
	}()

	go func(c chan os.Signal) {
		for {
			select {
			case <-c:
				return
			case <-time.After(time.Minute * 5):
				err := filepath.Walk(opt.DataPath, func(path string, info fs.FileInfo, err error) error {
					if info == nil {
						return nil
					}

					if info.IsDir() {
						return nil
					}

					switch filepath.Ext(path) {
					case ".bad":
						log.Warnf("remove bad file:%s", path)
						err := os.RemoveAll(path)
						if err != nil {
							log.Errorf("err:%v", err)
						}
					case ".tmp":
						if time.Since(info.ModTime()) > xtime.Day {
							log.Warnf("remove bad file:%s", path)
							err := os.RemoveAll(path)
							if err != nil {
								log.Errorf("err:%v", err)
							}
						}
					}
					return nil
				})
				if err != nil {
					log.Errorf("err:%v", err)
				}
			}
		}
	}(utils.GetExitSign())

	return nil
}

func Close() {
	for topicName, topic := range producer.CloneTopic() {
		for channelName, channel := range topic.CloneChannel() {
			err := channel.Close()
			if err != nil {
				log.Errorf("close channel %s-%s err:%v", topicName, channelName, err)
			}
		}

		err := topic.Close()
		if err != nil {
			log.Errorf("close topic %s err:%v", topicName, err)
		}
	}
}
