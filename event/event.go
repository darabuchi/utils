package event

import (
	"bytes"
	"os"
	"sync"
	"time"

	"github.com/darabuchi/log"
	"github.com/darabuchi/utils"
)

const worker = 3

type Event struct {
	Name string
	Data *EventData
}

type eventMsg struct {
	Name    string
	Data    *EventData
	TraceId string
}

var lock sync.RWMutex
var eventMap map[string][]Listener

var c chan eventMsg

func init() {
	eventMap = make(map[string][]Listener)
	c = make(chan eventMsg, 10)

	for i := 0; i < worker; i++ {
		go func(sign chan os.Signal) {
			for {
				select {
				case msg := <-c:
					lock.RLock()
					e := eventMap[msg.Name]
					lock.RUnlock()

					var w sync.WaitGroup
					for _, listener := range e {
						w.Add(1)
						go func(listener Listener) {
							defer w.Done()
							defer log.DelTrace()
							defer utils.CachePanic()
							traceId := bytes.NewBufferString(msg.TraceId)
							if msg.TraceId != "" {
								traceId.WriteString(".")
							}

							traceId.WriteString(log.GenTraceId())

							log.SetTrace(traceId.String())

							listener(Event{
								Name: msg.Name,
								Data: msg.Data,
							})
						}(listener)
					}
					w.Wait()
				case <-sign:
					return
				}
			}
		}(utils.GetExitSign())
	}
}

type Listener func(event Event)

func On(name string, listener Listener) {
	lock.Lock()
	defer lock.Unlock()
	eventMap[name] = append(eventMap[name], listener)
}

func Close(name string) {
	lock.Lock()
	defer lock.Unlock()
	delete(eventMap, name)
}

func Trigger(name string, data *EventData) {
	log.Infof("send event to %s", name)
	c <- eventMsg{
		Name:    name,
		Data:    data,
		TraceId: log.GetTrace(),
	}
}

func TriggerWithTimeout(name string, data *EventData, timeout time.Duration) {
	log.Infof("send event to %s", name)

	ticker := time.NewTimer(timeout)
	defer ticker.Stop()
	select {
	case c <- eventMsg{
		Name:    name,
		Data:    data,
		TraceId: log.GetTrace(),
	}:
	// 正常分支，不做特殊处理
	case <-ticker.C:
		log.Warnf("send event %s timeout", name)
	}
}
