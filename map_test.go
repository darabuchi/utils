package utils

import (
	"reflect"
	"testing"

	"github.com/darabuchi/log"
)

func TestNewMap(t *testing.T) {
	m := NewMap(map[string]interface{}{
		"a": []uint64{1, 2, 3, 4, 5},
	}).EnableCut(".")

	// log.Info(m.Get("a.b"))
	// log.Info(m.Get("a.d"))
	log.Info(m.GetSlice("a"))
}

func TestNewJson(t *testing.T) {
	m, err := NewMapWithJson([]byte(`{
    "code": 0,
    "data": {
        "request_id": "3d29133b-e74e-40c7-8274-46da235aacd0"
    },
    "msg": "suc"
}`))
	if err != nil {
		t.Errorf("err:%v", err)
		return
	}

	m.EnableCut(".")

	mm := m.ToMap()
	t.Log(MapKeysString(mm))
	data := mm["data"]
	t.Logf("%+v", data)
	requestId := data.(map[string]interface{})["request_id"]
	t.Log(requestId)
	t.Log(reflect.TypeOf(requestId).Kind())
}
