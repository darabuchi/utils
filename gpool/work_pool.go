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

	workerLock, poolLock sync.RWMutex

	worker, maxWorker *atomic.Uint32

	poolMap map[string]*Pool

	// 各种消息通道
	subWorkerCloseChan chan bool
}

type poolQueue struct {
	logic func(i interface{})
	args  interface{}
}

func NewPool(maxWorker int) *workPool {
	p := &workPool{
		maxWorker: atomic.NewUint32(uint32(maxWorker)),
		poolMap:   map[string]*Pool{},
		worker:    atomic.NewUint32(0),

		subWorkerCloseChan: make(chan bool, 100),
	}

	go func(sign chan os.Signal) {
		defer log.DelTrace()
		for {
			select {
			case <-time.After(time.Second * 30):
				// 定时刷新的兜底逻辑

				log.SetTrace(fmt.Sprintf("work_pool.check.%s", sandid.New().String()))
				p.checkPools()
			case <-p.subWorkerCloseChan:
				//  有新的资源释放的时候的优化逻辑
				log.SetTrace(fmt.Sprintf("work_pool.on_work_free.%s", sandid.New().String()))
				for {
					// 把多余的一起取完，不要处理太多次
					select {
					case <-p.subWorkerCloseChan:
					case <-time.After(time.Millisecond * 1500):
						goto END
					}
				}

			END:
				p.checkPools()

			case <-sign:
				log.Info("exist pool")
				return
			}
		}
	}(utils.GetExitSign())

	return p
}

func (p *workPool) checkPools() {
	poolMap := p.clonePool()
	var totalTask uint64
	for _, w := range poolMap {
		if w == nil {
			continue
		}

		if w.needClose() {
			w.Close()
			continue
		} else if w.needMoreWorker() {
			log.Infof("%s(%s) need new worker", w.name, w.id)
			w.tryApply()
		}

		log.Infof("%s(%s) worker:%d|max:%d|task total:%d|total:%d|wait:%d",
			w.name, w.id, w.worker.Load(), w.maxWorker.Load(), w.taskTotal.Load(), w.total.Load(), len(w.wait))
		totalTask += w.taskTotal.Load()
	}

	log.Infof("worker pool:%d|worker:%d|max:%d|total:%d", len(poolMap), p.worker.Load(), p.maxWorker.Load(), totalTask)
}

func (p *workPool) SetMaxWorker(worker int) {
	p.maxWorker.Store(uint32(worker))
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

func (p *workPool) LoadTotal() uint64 {
	var totalTask uint64
	for _, w := range p.clonePool() {
		if w == nil {
			continue
		}

		totalTask += w.taskTotal.Load()
	}
	return totalTask
}

func (p *workPool) Close() {
	for _, w := range p.clonePool() {
		if w == nil {
			continue
		}
		log.Infof("close pool %s(%s)", w.name, w.id)
		w.Close()
	}

	close(p.subWorkerCloseChan)
}

func (p *workPool) onWorkerFree() {
	p.subWorkerCloseChan <- true
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

func SetPoolGlobalMaxWorker(worker int) {
	_pool.SetMaxWorker(worker)
}
