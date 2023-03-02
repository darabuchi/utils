package utils_test

import (
	"github.com/darabuchi/log"
	"github.com/darabuchi/utils"
	"reflect"
	"testing"
)

func TestNewMap(t *testing.T) {
	m := utils.NewMap(map[string]interface{}{
		"a": []uint64{1, 2, 3, 4, 5},
	}).EnableCut(".")

	// log.Info(m.Get("a.b"))
	// log.Info(m.Get("a.d"))
	log.Info(m.GetSlice("a"))
	log.Info(utils.ToString(m.ToMap()))
}

func TestNewJson(t *testing.T) {
	m, err := utils.NewMapWithJson([]byte(`{
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
	t.Log(utils.MapKeysString(mm))
	data := mm["data"]
	t.Logf("%+v", data)
	requestId := data.(map[string]interface{})["request_id"]
	t.Log(requestId)
	t.Log(reflect.TypeOf(requestId).Kind())
}

func TestBeginCut(t *testing.T) {
	m := utils.NewMap(map[string]interface{}{
		".a":   1,
		".b.a": 2,
	})
	m.EnableCut(".")

	t.Log(m.Get(".a"))
}
