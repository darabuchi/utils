package unit_test

import (
	"github.com/darabuchi/utils/unit"
	"testing"
)

func TestFormatInt64(t *testing.T) {
	t.Log(unit.FormatInt64(111111111, 2))
}
