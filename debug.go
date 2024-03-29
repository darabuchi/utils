package utils

import (
	"runtime/debug"
	"strings"

	"github.com/darabuchi/log"
)

func CachePanic() {
	CachePanicWithHandle(nil)
}

func CachePanicWithHandle(handle func(err interface{})) {
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
		if handle != nil {
			handle(err)
		}
	}
}

func PrintStack() {
	st := debug.Stack()
	if len(st) > 0 {
		log.Info("dump stack:")
		log.Info(string(st))
	} else {
		log.Info("stack is empty")
	}
}
