package utils_test

import (
	"testing"

	"github.com/darabuchi/utils"
	"github.com/darabuchi/utils/xtime"
)

func TestDuration(t *testing.T) {
	t.Log(utils.FormatDuration(xtime.Year * 10))
	t.Log(utils.FormatDuration(xtime.Month * 10))
	t.Log(utils.FormatDuration(xtime.Day * 10))
	t.Log(utils.FormatDuration(xtime.Hour * 10))
	t.Log(utils.FormatDuration(xtime.Minute * 10))
	t.Log(utils.FormatDuration(xtime.Second * 10))
}
