package utils

import (
	"testing"

	"github.com/darabuchi/log"
)

func TestIsLocalIp(t *testing.T) {
	log.Info(IsLocalIp("127.0.0.1"))
}
