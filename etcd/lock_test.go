package etcd

import (
	"sync"
	"testing"
	"time"

	"github.com/darabuchi/log"
)

func TestLock(t *testing.T) {
	err := Connect(Config{
		Addrs: []string{"http://127.0.0.1:2379"},
	})
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}

	var w sync.WaitGroup
	logic := func(i int) {
		defer w.Done()

		lock := NewMutex("test", time.Second*5)

		if i%2 == 0 {
			err := lock.LockWithTimeout(time.Second * 10)
			if err != nil {
				log.Errorf("err:%v", err)
				return
			}
		} else {
			err := lock.Lock()
			if err != nil {
				log.Errorf("err:%v", err)
				return
			}
		}

		log.Info(i)
		time.Sleep(time.Second * 3)

		err = lock.Unlock()
		if err != nil {
			log.Errorf("err:%v", err)
			return
		}
	}

	for i := 0; i < 100; i++ {
		w.Add(1)
		go logic(i)
	}

	w.Wait()
}
