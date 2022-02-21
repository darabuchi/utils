package event

import (
	"sync"
)

type EventData struct {
	lock sync.RWMutex

	data map[string]interface{}
}

func NewData(data map[string]interface{}) *EventData {
	return &EventData{
		data: data,
	}
}

func (p *EventData) Set(key string, val interface{}) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.data[key] = val
}

func (p *EventData) Get(key string) interface{} {
	p.lock.RLock()
	defer p.lock.RUnlock()
	return p.data[key]
}

func (p *EventData) GetDef(key string, def interface{}) interface{} {
	p.lock.RLock()
	defer p.lock.RUnlock()
	if val, ok := p.data[key]; ok {
		return val
	}
	return def
}

func (p *EventData) GetAll() map[string]interface{} {
	p.lock.RLock()
	defer p.lock.RUnlock()
	return p.data
}

func (p *EventData) GetInt32(key string) int32 {
	val := p.data[key]
	if val == nil {
		return 0
	}

	return val.(int32)
}

func (p *EventData) GetInt64(key string) int64 {
	val := p.data[key]
	if val == nil {
		return 0
	}

	return val.(int64)
}

func (p *EventData) GetUint32(key string) uint32 {
	val := p.data[key]
	if val == nil {
		return 0
	}

	return val.(uint32)
}

func (p *EventData) GetUint64(key string) uint64 {
	val := p.data[key]
	if val == nil {
		return 0
	}

	return val.(uint64)
}

func (p *EventData) GetStr(key string) string {
	val := p.data[key]
	if val == nil {
		return ""
	}

	return val.(string)
}

func (p *EventData) GetBool(key string) bool {
	val := p.data[key]
	if val == nil {
		return false
	}

	return val.(bool)
}
