package gpool

import (
	"fmt"
	"github.com/aofei/sandid"
	"go.uber.org/atomic"
	"sync"
)

type SubPool struct {
	id   string
	name string

	defaultLogic Logic

	w     sync.WaitGroup
	total *atomic.Int64

	pool *Pool

	lock sync.RWMutex

	stop bool
}

func newSubPoolWithFunc(name string, pool *Pool, logic Logic) *SubPool {
	s := SubPool{
		name:         fmt.Sprintf("%s(%s)-%s", pool.name, pool.id, name),
		id:           sandid.New().String(),
		defaultLogic: logic,
		total:        atomic.NewInt64(0),
		pool:         pool,
	}

	return &s
}

func (p *SubPool) Submit(i interface{}) {
	p.SubmitWithFunc(i, p.defaultLogic)
}

func (p *SubPool) SetLogic(logic Logic) {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.defaultLogic = logic
}

func (p *SubPool) SubmitWithFunc(i interface{}, logic Logic) {
	p.w.Add(1)
	p.total.Inc()
	p.pool.SubmitWithFunc(i, func(i interface{}) {
		defer func() {
			p.total.Dec()
			p.w.Done()
		}()

		if p.checkStop() {
			return
		}

		logic(i)
	})
}

func (p *SubPool) SubmitWait(i interface{}) {
	p.SubmitWithFuncWait(i, p.defaultLogic)
}

func (p *SubPool) SubmitWithFuncWait(i interface{}, logic Logic) {
	c := make(chan bool)
	defer close(c)
	p.SubmitWithFunc(i, func(i interface{}) {
		logic(i)
		c <- true
	})
	<-c
}

func (p *SubPool) Wait() {
	p.w.Wait()
}

func (p *SubPool) Close() {
	p.lock.Lock()
	defer p.lock.Unlock()
	p.stop = true
}

func (p *SubPool) checkStop() bool {
	p.lock.RLock()
	defer p.lock.RUnlock()
	return p.stop
}

func (p *SubPool) LoadWaitCount() int64 {
	return p.total.Load()
}
