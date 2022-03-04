//go:build windows
// +build windows

package utils

import (
	"os"
	"syscall"
)

var exitSignal = []os.Signal{
	os.Interrupt,
	syscall.SIGHUP,
	syscall.SIGINT,
	syscall.SIGQUIT,
	syscall.SIGILL,
	syscall.SIGABRT,
	syscall.SIGKILL,
	syscall.SIGSEGV,
	syscall.SIGTERM,
}
