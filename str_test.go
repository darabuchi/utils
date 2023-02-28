package utils_test

import (
	"github.com/darabuchi/utils"
	"testing"
)

func TestSnakeToCamel(t *testing.T) {
	t.Log(utils.SnakeToCamel("user_hook"))
}
