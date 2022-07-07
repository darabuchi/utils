package mq

import (
	"fmt"

	"github.com/nsqio/nsq/nsqd"
)

func CloneTopic() map[string]*nsqd.Topic {
	return producer.CloneTopic()
}

func GetTopic(name fmt.Stringer) *nsqd.Topic {
	return producer.GetTopic(name.String())
}

func GetTopicDepth(name fmt.Stringer) int64 {
	topic := GetTopic(name)

	var depth int64
	for _, channel := range topic.CloneChannel() {
		depth += channel.Depth() + channel.InFlightDepth() + channel.DeferredDepth()
	}

	return depth
}
