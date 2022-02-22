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

	SetPoolGlobalMaxWorker(3)

	var w sync.WaitGroup

	total := atomic.NewInt64(0)

	newPool := func() *Pool {
		total.Inc()
		return NewPoolGlobalWithFunc(total.String(), 2, func(i interface{}) {
			defer w.Done()
			time.Sleep(time.Duration(rand.Intn(3)+i.(int)) % 3 * time.Second)
		}).SetTimeout(time.Second * 10)
	}

	submit := func(pool *Pool, max int) {
		for i := 0; i < max; i++ {
			w.Add(1)
			pool.Submit(i + 1)
		}
	}

	run := func(max int) {
		p := newPool()
		submit(p, 25)
	}

	for i := 0; i < 25; i++ {
		run(10)
	}

	log.Infof("wait finish")
	w.Wait()
}

func TestSubPool(t *testing.T) {
	log.SetLevel(log.InfoLevel)

	SetPoolGlobalMaxWorker(3)

	var w sync.WaitGroup
	pool := NewPoolGlobalWithFunc("", 3, func(i interface{}) {
		defer w.Done()
		time.Sleep(time.Duration(rand.Intn(3)+i.(int)) % 3 * time.Second)
	}).SetTimeout(time.Second * 10)

	total := atomic.NewInt64(0)
	newPool := func() *SubPool {
		total.Inc()
		return pool.NewSubPool(total.String())
	}

	submit := func(pool *SubPool, max int) {
		for i := 0; i < max; i++ {
			w.Add(1)
			pool.Submit(i + 1)
		}
	}

	run := func(max int) {
		p := newPool()
		submit(p, 25)
	}

	for i := 0; i < 25; i++ {
		run(10)
	}

	log.Infof("wait finish")
	w.Wait()
}
