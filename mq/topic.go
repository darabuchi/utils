package mq

import (
	"github.com/nsqio/nsq/nsqd"
)

func CloneTopic() map[string]*nsqd.Topic {
	return producer.CloneTopic()
}

func GetTopic(name string) *nsqd.Topic {
	return producer.GetTopic(name)
}

func GetTopicDepth(name string) int64 {
	topic := GetTopic(name)
	
	var depth int64
	for _, channel := range topic.CloneChannel() {
		depth += channel.Depth() + channel.InFlightDepth() + channel.DeferredDepth()
	}
	
	return depth
}
