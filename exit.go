package utils

import (
	"os"
	"os/signal"
	"syscall"
)

func GetExitSign() chan os.Signal {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGQUIT)
	return sigCh
}
