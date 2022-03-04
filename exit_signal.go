//go:build !plan9 && !windows
// +build !plan9 && !windows

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
	syscall.SIGCONT,
	syscall.SIGSTOP,
	syscall.SIGTSTP,
}
