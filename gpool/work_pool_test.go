package gpool

import (
	"github.com/darabuchi/log"
	"go.uber.org/atomic"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestNewSubPool(t *testing.T) {
	total := atomic.NewUint64(0)
	handled := atomic.NewUint64(0)
	pool := NewPoolGlobal("test_pool", 5)
	pool.SetDefLogic(func(i interface{}) {
		time.Sleep(time.Second)
		log.Infof("total:%d|handled:%d|work:%d|wait:%d", total.Load(), handled.Inc(), pool.worker.Load(), len(pool.wait))
	})
	defer pool.Close()

	for j := 0; j < 20; j++ {
		for i := 0; i < 100; i++ {
			total.Inc()
			pool.Submit(nil)
		}
		time.Sleep(time.Second * 3)
	}

	pool.Wait()
}

func TestPools(t *testing.T) {
	log.SetLevel(log.InfoLevel)
	var w sync.WaitGroup

	total := atomic.NewInt64(0)

	newPool := func() *Pool {
		total.Inc()
		return NewPoolGlobalWithFunc(total.String(), 5, func(i interface{}) {
			defer w.Done()
			time.Sleep(time.Duration(rand.Intn(3)+i.(int)) % 10 * time.Second)
		})
	}

	submit := func(pool *Pool, max int) {
		for i := 0; i < max; i++ {
			w.Add(1)
			pool.Submit(i + 1)
		}
	}

	run := func(max int) {
		p := newPool()

		w.Add(1)

		go func() {
			log.SetTrace("")
			defer log.DelTrace()

			defer w.Done()
			for i := 0; i < max; i++ {
				submit(p, rand.Intn(100))
				time.Sleep(time.Duration(rand.Intn(5)+i+1) * time.Second)
			}
		}()
	}

	for i := 0; i < 50; i++ {
		run(rand.Intn(20) + 1)
	}

	w.Wait()
}
