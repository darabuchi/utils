package utils

import (
	"github.com/darabuchi/log"
	"testing"
)

func TestInfo(t *testing.T) {
	log.Info("msg")
	//log.SetModule("model")
	log.Info("msg")
	log.Infof("%s", "msg")
}
