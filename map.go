package utils

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"

	"go.uber.org/atomic"
	"gopkg.in/yaml.v3"
)

type Map struct {
	data *sync.Map
	cut  *atomic.Bool
	seq  *atomic.String
}

var (
	ErrNotFound = errors.New("not found")
)

func NewMap(m map[string]interface{}) *Map {
	m2 := &Map{
		data: &sync.Map{},
		cut:  atomic.NewBool(false),
		seq:  atomic.NewString(""),
	}
	for k, v := range m {
		m2.data.Store(k, v)
	}
	return m2
}

func NewMapWithJson(s []byte) (*Map, error) {
	var m map[string]interface{}
	err := json.Unmarshal(s, &m)
	if err != nil {
		return nil, err
	}
	return NewMap(m), nil
}

func NewMapWithYaml(s []byte) (*Map, error) {
	var m map[string]interface{}
	err := yaml.Unmarshal(s, &m)
	if err != nil {
		return nil, err
	}
	return NewMap(m), nil
}

func NewMapWithAny(s interface{}) (*Map, error) {
	buf, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	var m map[string]interface{}
	err = yaml.Unmarshal(buf, &m)
	if err != nil {
		return nil, err
	}
	return NewMap(m), nil
}

func (p *Map) EnableCut(seq string) *Map {
	p.cut.Store(true)
	p.seq.Store(seq)
	return p
}

func (p *Map) DisableCut() *Map {
	p.cut.Store(false)
	return p
}

func (p *Map) Set(key string, value interface{}) {
	p.data.Store(key, value)
}

func (p *Map) Get(key string) (interface{}, error) {
	return p.get(key)
}

func (p *Map) get(key string) (interface{}, error) {
	var val interface{}
	var ok bool
	if !p.cut.Load() {
		if val, ok := p.data.Load(key); ok {
			return val, nil
		}
		return nil, ErrNotFound
	}

	seq := p.seq.Load()
	keys := strings.Split(key, seq)

	data := p.data
	var m *Map
	for len(keys) > 1 {
		k := keys[0]
		keys = keys[1:]

		val, ok = data.Load(k)
		if !ok {
			return nil, ErrNotFound
		}

		m = p.toMap(val)
		if m == nil {
			return nil, ErrNotFound
		}

		data = m.data
	}

	if len(keys) > 0 {
		if val, ok = data.Load(keys[0]); ok {
			return val, nil
		}
		return nil, ErrNotFound
	}

	return nil, ErrNotFound
}

func (p *Map) Exists(key string) bool {
	if _, ok := p.data.Load(key); ok {
		return true
	}
	return false
}

func (p *Map) GetBool(key string) bool {
	val, err := p.get(key)
	if err != nil {
		return false
	}

	switch x := val.(type) {
	case bool:
		return x
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64:
		return x != 0
	case string:
		switch strings.ToLower(x) {
		case "true", "1":
			return true
		case "false", "0":
			return false
		case "":
			return false
		default:
			return true
		}
	case []byte:
		switch string(bytes.ToLower(x)) {
		case "true", "1":
			return true
		case "false", "0":
			return false
		case "":
			return false
		default:
			return true
		}
	default:
		return val == nil
	}
}

func (p *Map) GetInt32(key string) int32 {
	val, err := p.get(key)
	if err != nil {
		return 0
	}

	switch x := val.(type) {
	case bool:
		if x {
			return 1
		}
		return 0
	case int:
		return int32(x)
	case int8:
		return int32(x)
	case int16:
		return int32(x)
	case int32:
		return x
	case int64:
		return int32(x)
	case uint:
		return int32(x)
	case uint8:
		return int32(x)
	case uint16:
		return int32(x)
	case uint32:
		return int32(x)
	case uint64:
		return int32(x)
	case float32:
		return int32(x)
	case float64:
		return int32(x)
	case string:
		val, err := strconv.ParseUint(x, 10, 16)
		if err != nil {
			return 0
		}
		return int32(val)
	case []byte:
		val, err := strconv.ParseUint(string(x), 10, 16)
		if err != nil {
			return 0
		}
		return int32(val)
	default:
		return 0
	}
}

func (p *Map) GetUint16(key string) uint16 {
	val, err := p.get(key)
	if err != nil {
		return 0
	}

	switch x := val.(type) {
	case bool:
		if x {
			return 1
		}
		return 0
	case int:
		return uint16(x)
	case int8:
		return uint16(x)
	case int16:
		return uint16(x)
	case int32:
		return uint16(x)
	case int64:
		return uint16(x)
	case uint:
		return uint16(x)
	case uint8:
		return uint16(x)
	case uint16:
		return x
	case uint32:
		return uint16(x)
	case uint64:
		return uint16(x)
	case float32:
		return uint16(x)
	case float64:
		return uint16(x)
	case string:
		val, err := strconv.ParseUint(x, 10, 16)
		if err != nil {
			return 0
		}
		return uint16(val)
	case []byte:
		val, err := strconv.ParseUint(string(x), 10, 16)
		if err != nil {
			return 0
		}
		return uint16(val)
	default:
		return 0
	}
}

func (p *Map) GetUint32(key string) uint32 {
	val, err := p.get(key)
	if err != nil {
		return 0
	}

	switch x := val.(type) {
	case bool:
		if x {
			return 1
		}
		return 0
	case int:
		return uint32(x)
	case int8:
		return uint32(x)
	case int16:
		return uint32(x)
	case int32:
		return uint32(x)
	case int64:
		return uint32(x)
	case uint:
		return uint32(x)
	case uint8:
		return uint32(x)
	case uint16:
		return uint32(x)
	case uint32:
		return x
	case uint64:
		return uint32(x)
	case float32:
		return uint32(x)
	case float64:
		return uint32(x)
	case string:
		val, err := strconv.ParseUint(x, 10, 16)
		if err != nil {
			return 0
		}
		return uint32(val)
	case []byte:
		val, err := strconv.ParseUint(string(x), 10, 16)
		if err != nil {
			return 0
		}
		return uint32(val)
	default:
		return 0
	}
}

func (p *Map) GetUint64(key string) uint64 {
	val, err := p.get(key)
	if err != nil {
		return 0
	}

	switch x := val.(type) {
	case bool:
		if x {
			return 1
		}
		return 0
	case int:
		return uint64(x)
	case int8:
		return uint64(x)
	case int16:
		return uint64(x)
	case int32:
		return uint64(x)
	case int64:
		return uint64(x)
	case uint:
		return uint64(x)
	case uint8:
		return uint64(x)
	case uint16:
		return uint64(x)
	case uint32:
		return uint64(x)
	case uint64:
		return x
	case float32:
		return uint64(x)
	case float64:
		return uint64(x)
	case string:
		val, err := strconv.ParseUint(x, 10, 16)
		if err != nil {
			return 0
		}
		return uint64(val)
	case []byte:
		val, err := strconv.ParseUint(string(x), 10, 16)
		if err != nil {
			return 0
		}
		return uint64(val)
	default:
		return 0
	}
}

func (p *Map) GetFloat64(key string) float64 {
	val, err := p.get(key)
	if err != nil {
		return 0
	}

	switch x := val.(type) {
	case bool:
		if x {
			return 1
		}
		return 0
	case int:
		return float64(x)
	case int8:
		return float64(x)
	case int16:
		return float64(x)
	case int32:
		return float64(x)
	case int64:
		return float64(x)
	case uint:
		return float64(x)
	case uint8:
		return float64(x)
	case uint16:
		return float64(x)
	case uint32:
		return float64(x)
	case uint64:
		return float64(x)
	case float32:
		return float64(x)
	case float64:
		return x
	case string:
		val, err := strconv.ParseUint(x, 10, 16)
		if err != nil {
			return 0
		}
		return float64(val)
	case []byte:
		val, err := strconv.ParseFloat(string(x), 64)
		if err != nil {
			return 0
		}
		return val
	default:
		return 0
	}
}

func (p *Map) GetString(key string) string {
	val, err := p.get(key)
	if err != nil {
		return ""
	}

	return p.toString(val)
}

func (p *Map) toString(val interface{}) string {
	switch x := val.(type) {
	case bool:
		if x {
			return "1"
		}
		return "0"
	case int:
		return fmt.Sprintf("%d", x)
	case int8:
		return fmt.Sprintf("%d", x)
	case int16:
		return fmt.Sprintf("%d", x)
	case int32:
		return fmt.Sprintf("%d", x)
	case int64:
		return fmt.Sprintf("%d", x)
	case uint:
		return fmt.Sprintf("%d", x)
	case uint8:
		return fmt.Sprintf("%d", x)
	case uint16:
		return fmt.Sprintf("%d", x)
	case uint32:
		return fmt.Sprintf("%d", x)
	case uint64:
		return fmt.Sprintf("%d", x)
	case float32:
		if math.Floor(float64(x)) == float64(x) {
			return fmt.Sprintf("%.0f", x)
		}

		return fmt.Sprintf("%f", x)
	case float64:
		if math.Floor(x) == x {
			return fmt.Sprintf("%.0f", x)
		}

		return fmt.Sprintf("%f", x)
	case string:
		return x
	case []byte:
		return string(x)
	default:
		return ""
	}
}

func (p *Map) GetBytes(key string) []byte {
	val, err := p.get(key)
	if err != nil {
		return []byte("")
	}

	switch x := val.(type) {
	case bool:
		if x {
			return []byte("1")
		}
		return []byte("0")
	case int:
		return []byte(fmt.Sprintf("%d", x))
	case int8:
		return []byte(fmt.Sprintf("%d", x))
	case int16:
		return []byte(fmt.Sprintf("%d", x))
	case int32:
		return []byte(fmt.Sprintf("%d", x))
	case int64:
		return []byte(fmt.Sprintf("%d", x))
	case uint:
		return []byte(fmt.Sprintf("%d", x))
	case uint8:
		return []byte(fmt.Sprintf("%d", x))
	case uint16:
		return []byte(fmt.Sprintf("%d", x))
	case uint32:
		return []byte(fmt.Sprintf("%d", x))
	case uint64:
		return []byte(fmt.Sprintf("%d", x))
	case float32:
		return []byte(fmt.Sprintf("%v", x))
	case float64:
		return []byte(fmt.Sprintf("%v", x))
	case string:
		return []byte(x)
	case []byte:
		return x
	default:
		return []byte("")
	}
}

func (p *Map) GetMap(key string) *Map {
	val, err := p.get(key)
	if err != nil {
		return NewMap(nil)
	}

	return p.toMap(val)
}

func (p *Map) toMap(val interface{}) *Map {
	switch x := val.(type) {
	case bool, int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64:
		return NewMap(nil)
	case string:
		var m map[string]interface{}
		err := json.Unmarshal([]byte(x), &m)
		if err != nil {
			return NewMap(nil)
		}
		return NewMap(m)
	case []byte:
		var m map[string]interface{}
		err := json.Unmarshal(x, &m)
		if err != nil {
			return NewMap(nil)
		}
		return NewMap(m)
	case map[string]interface{}:
		return NewMap(x)
	case map[interface{}]interface{}:
		m := NewMap(nil)
		for k, v := range x {
			m.Set(p.toString(k), v)
		}
		return m
	default:
		buf, err := json.Marshal(x)
		if err != nil {
			return NewMap(nil)
		}
		var m map[string]interface{}
		err = json.Unmarshal(buf, &m)
		if err != nil {
			return NewMap(nil)
		}
		return NewMap(m)
	}
}

func (p *Map) ToSyncMap() *sync.Map {
	var m sync.Map
	p.data.Range(func(key, value interface{}) bool {
		m.Store(key, value)
		return true
	})
	return &m
}

func (p *Map) ToMap() map[string]interface{} {
	m := map[string]interface{}{}
	p.data.Range(func(key, value interface{}) bool {
		m[key.(string)] = value
		return true
	})
	return m
}

func (p *Map) Clone() *Map {
	return &Map{
		data: p.ToSyncMap(),
		cut:  atomic.NewBool(p.cut.Load()),
		seq:  atomic.NewString(p.seq.Load()),
	}
}