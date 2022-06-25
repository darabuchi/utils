//go:build !plan9 && !windows && !darwin && !dragonfly && !freebsd && !linux && !solaris && !openbsd && !netbsd
// +build !plan9,!windows,!darwin,!dragonfly,!freebsd,!linux,!solaris,!openbsd,!netbsd

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
	syscall.SIGPWR,
	syscall.SIGPIPE,
	syscall.SIGSTKFLT,
	syscall.SIGFPE,
}
