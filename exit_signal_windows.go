//go:build windows
// +build windows

package utils

import (
	"os"
	"syscall"
)

var exitSignal = []os.Signal{
	os.Interrupt,
	os.Interrupt,
	syscall.SIGHUP,
	syscall.SIGINT,
	syscall.SIGQUIT,
	syscall.SIGILL,
	syscall.SIGABRT,
	syscall.SIGBUS,
	syscall.SIGKILL,
	syscall.SIGSEGV,
	syscall.SIGTERM,
	syscall.SIGTRAP,
	syscall.SIGPIPE,
	syscall.SIGFPE,
}
