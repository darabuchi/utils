package utils

import (
	"testing"
)

func TestProcessBarAscii(t *testing.T) {
	t.Log(ProcessBarAscii(236, 53, 30))
	t.Log(ProcessBarAscii(100, 13, 30))
	t.Log(ProcessBarAscii(100, 10, 30))
	t.Log(ProcessBarAscii(100, 200, 30))
	t.Log(ProcessBarAscii(100, 0, 30))
}
