package gpool

import (
	"fmt"
	"github.com/aofei/sandid"
	"github.com/darabuchi/log"
	"github.com/darabuchi/utils"
	"go.uber.org/atomic"
	"os"
	"sync"
	"time"
)

type Logic func(i interface{})

type workPool struct {
	defaultLogic Logic

	workTotal *atomic.Uint64

	workerLock, poolLock sync.RWMutex

	worker, maxWorker *atomic.Uint32

	poolMap map[string]*Pool
}

type poolQueue struct {
	logic func(i interface{})
	args  interface{}
}

func NewPool(maxWorker int) *workPool {
	p := &workPool{
		workTotal: atomic.NewUint64(0),
		maxWorker: atomic.NewUint32(uint32(maxWorker)),
		poolMap:   map[string]*Pool{},
		worker:    atomic.NewUint32(0),
	}

	go func(sign chan os.Signal) {
		defer log.DelTrace()
		for {
			select {
			case <-time.After(time.Second * 30):
				log.SetTrace(fmt.Sprintf("work_pool.check.%s", sandid.New().String()))
				for _, w := range p.clonePool() {
					if w == nil {
						continue
					}

					if w.needClose() {
						w.Close()
					} else if w.needMoreWorker() {
						log.Infof("%s(%s) need new worker", w.name, w.id)
						w.tryApply()
					}

					log.Infof("%s(%s) worker:%d|max:%d|task total:%d|total:%d|wait:%d",
						w.name, w.id, w.worker.Load(), w.maxWorker.Load(), w.taskTotal.Load(), w.total.Load(), len(w.wait))
				}

				log.Infof("worker pool worker:%d|max:%d|total:%d", p.worker.Load(), p.maxWorker.Load(), p.workTotal.Load())
			case <-sign:
				log.Info("exist pool")
				return
			}
		}
	}(utils.GetExitSign())

	return p
}

func (p *workPool) clonePool() []*Pool {
	var l []*Pool

	p.poolLock.RLock()
	defer p.poolLock.RUnlock()
	for _, w := range p.poolMap {
		l = append(l, w)
	}

	return l
}

func (p *workPool) NewPool(name string, work int) *Pool {
	return p.NewPoolWithFunc(name, work, nil)
}

func (p *workPool) NewPoolWithFunc(name string, work int, logic Logic) *Pool {
	id := sandid.New().String()

	pool := newPool(name, id, p)
	pool.SetWorker(work)
	pool.SetDefLogic(logic)

	p.poolLock.Lock()
	p.poolMap[id] = pool
	p.poolLock.Unlock()

	log.Infof("new pool %s(%s)", name, id)

	return pool
}

func (p *workPool) freePool(id string) {
	p.poolLock.Lock()
	delete(p.poolMap, id)
	p.poolLock.Unlock()
}

func (p *workPool) applyWorker() bool {
	p.workerLock.Lock()
	defer p.workerLock.Unlock()

	if p.worker.Load() >= p.maxWorker.Load() {
		return false
	}

	log.Infof("now has %d worker pool", p.worker.Inc())

	return true
}

func (p *workPool) freeWorker() {
	p.worker.Dec()
}

func (p *workPool) inc() {
	p.workTotal.Inc()
}

func (p *workPool) dec() {
	p.workTotal.Dec()
}

func (p *workPool) done(j uint64) {
	p.workTotal.Sub(j)
}

func (p *workPool) LoadTotal() uint64 {
	return p.workTotal.Load()
}

func (p *workPool) Close() {
	for _, w := range p.clonePool() {
		if w == nil {
			continue
		}
		log.Infof("close pool %s(%s)", w.name, w.id)
		w.Close()
	}
}

var _pool *workPool

func init() {
	_pool = NewPool(32)
}

func NewPoolGlobal(name string, work int) *Pool {
	return NewPoolGlobalWithFunc(name, work, nil)
}

func NewPoolGlobalWithFunc(name string, work int, logic Logic) *Pool {
	return _pool.NewPoolWithFunc(name, work, logic)
}

func PoolGlobalLoadTotal() uint64 {
	return _pool.LoadTotal()
}
