package mq

import (
	"fmt"

	"github.com/nsqio/nsq/nsqd"
)

type Topic string

func (p Topic) String() string {
	return string(p)
}

func (p Topic) GoString() string {
	return p.String()
}

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

func GetDepth() (depth int64) {
	for _, topic := range CloneTopic() {
		for _, channel := range topic.CloneChannel() {
			depth += channel.Depth() + channel.InFlightDepth() + channel.DeferredDepth()
		}
	}
	return
}
