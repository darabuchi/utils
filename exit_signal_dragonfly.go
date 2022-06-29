//go:build dragonfly
// +build dragonfly

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
	syscall.SIGBUS,
	syscall.SIGKILL,
	syscall.SIGSEGV,
	syscall.SIGSYS,
	syscall.SIGTERM,
	syscall.SIGCONT,
	syscall.SIGSTOP,
	syscall.SIGTRAP,
	syscall.SIGTSTP,
	syscall.SIGPIPE,
	syscall.SIGFPE,
}
