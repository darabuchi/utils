package utils

import (
	log "github.com/sirupsen/logrus"
	"runtime/debug"
	"strings"
)

func CachePanic() {
	if err := recover(); err != nil {
		log.Errorf("PROCESS PANIC: err %s", err)
		st := debug.Stack()
		if len(st) > 0 {
			log.Errorf("dump stack (%s):", err)
			lines := strings.Split(string(st), "\n")
			for _, line := range lines {
				log.Error("  ", line)
			}
		} else {
			log.Errorf("stack is empty (%s)", err)
		}
	}
}
