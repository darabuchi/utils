package utils

import (
	"os"
	"os/signal"

	"github.com/darabuchi/log"
)

func GetExitSign() chan os.Signal {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, exitSignal...)
	return sigCh
}

func WaitExist() {
	<-GetExitSign()
}

func Exit() {
	process, err := os.FindProcess(os.Getpid())
	if err != nil {
		log.Errorf("err:%v", err)
		os.Exit(0)
	} else {
		log.Infof("will stop process:%d", process.Pid)
		err = process.Signal(os.Kill)
		if err != nil {
			log.Errorf("err:%v", err)
			os.Exit(0)
		}
	}
}
