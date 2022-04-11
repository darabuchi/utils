package utils

import (
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
