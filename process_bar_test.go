package utils_test

import (
	"github.com/darabuchi/utils"
	"testing"
)

func TestProcessBarAscii(t *testing.T) {
	t.Log(utils.ProcessBarAscii(236, 53, 30))
	t.Log(utils.ProcessBarAscii(100, 13, 30))
	t.Log(utils.ProcessBarAscii(100, 10, 30))
	t.Log(utils.ProcessBarAscii(100, 200, 30))
	t.Log(utils.ProcessBarAscii(100, 0, 30))
}
