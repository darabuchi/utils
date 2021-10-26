package utils

import (
	"os"
	"os/signal"
)

func GetExitSign() chan os.Signal {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, exitSignal...)
	return sigCh
}
