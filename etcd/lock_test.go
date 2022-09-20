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

		lock := NewMutex("test")
		err := lock.Lock()
		if err != nil {
			log.Errorf("err:%v", err)
			return
		}

		// if i%2 == 0 {
		// 	return
		// }

		time.Sleep(time.Second * 20)
		log.Info(i)

		err = lock.Unlock()
		if err != nil {
			log.Errorf("err:%v", err)
			return
		}
	}

	for i := 0; i < 5; i++ {
		w.Add(1)
		go logic(i)
	}

	w.Wait()
}

func TestWhenLock(t *testing.T) {
	err := Connect(Config{
		Addrs: []string{"http://127.0.0.1:2379"},
	})
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}

	lock := NewMutex("test")
	err = lock.Lock()
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}

	err = lock.Lock()
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}

	log.Info(time.Now().String())

	err = lock.Lock()
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}

	log.Info(time.Now().String())

	err = lock.Lock()
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}

	log.Info(time.Now().String())
}
