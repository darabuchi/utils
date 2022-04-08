package utils

import (
	"testing"

	"github.com/darabuchi/log"
)

func TestNewMap(t *testing.T) {
	m := NewMap(map[string]interface{}{
		"a": map[string]interface{}{
			"b": 1,
			"c": map[string]interface{}{
				"d": float64(1.1),
			},
		},
	}).EnableCut(".")

	// log.Info(m.Get("a.b"))
	// log.Info(m.Get("a.d"))
	log.Info(m.GetString("a.c.d"))
}
