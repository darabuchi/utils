package cache

import (
	"testing"

	"github.com/darabuchi/log"
)

func TestInit(t *testing.T) {
	log.Info(Connect("127.0.0.1:6379", 0, ""))
	Set("test", 1)
	log.Info(Get("test"))
}
