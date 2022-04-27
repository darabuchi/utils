package gpool

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"time"

	"github.com/darabuchi/log"
	"github.com/darabuchi/utils"
	"go.uber.org/atomic"
)

type Logic func(i interface{})

type WorkPool struct {
	log *log.Logger

	defaultLogic Logic

	workerLock, poolLock sync.RWMutex

	worker, maxWorker *atomic.Uint32

	poolMap map[string]*Pool

	// 各种消息通道
	subWorkerCloseChan chan bool
}

type poolQueue struct {
	logic   func(i interface{})
	args    interface{}
	traceId string
}

func NewPool(maxWorker int) *WorkPool {
	p := &WorkPool{
		log:       log.Clone(),
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

				log.SetTrace(fmt.Sprintf("work_pool.check.%s", log.GenTraceId()))
				p.checkPools()
			case <-p.subWorkerCloseChan:
				//  有新的资源释放的时候的优化逻辑
				log.SetTrace(fmt.Sprintf("work_pool.on_work_free.%s", log.GenTraceId()))
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
				p.log.Info("exist pool")
				return
			}
		}
	}(utils.GetExitSign())

	return p
}

func (p *WorkPool) SetLogger(l *log.Logger) *WorkPool {
	p.log = l
	return p
}

func (p *WorkPool) checkPools() {
	poolMap := p.clonePool()
	var totalTask uint64
	for _, w := range poolMap {
		if w == nil {
			continue
		}

		if w.needClose() {
			w.Close()
			continue
		}

		if w.needMoreWorker() {
			p.log.Infof("%s(%s) need new worker", w.name, w.id)
			// 申请失败，并且需要强制申请
			if !w.tryApply() && w.needMoreWorkerForce() {
				w.applyForce()
			}
		}

		if p.worker.Load() >= p.maxWorker.Load() {
			if w.needMoreFree() {
				w.free()
			}
		}

		p.log.Infof("%s(%s) worker:%d|max:%d|task total:%d|wait:%d",
			w.name, w.id, w.worker.Load(), w.maxWorker.Load(), w.taskTotal.Load(), len(w.wait))
		totalTask += w.taskTotal.Load()
	}

	p.log.Infof("worker pool:%d|worker:%d|max:%d|total:%d", len(poolMap), p.worker.Load(), p.maxWorker.Load(), totalTask)
}

func (p *WorkPool) SetMaxWorker(worker int) *WorkPool {
	p.maxWorker.Store(uint32(worker))
	return p
}

func (p *WorkPool) clonePool() []*Pool {
	var l []*Pool

	p.poolLock.RLock()
	defer p.poolLock.RUnlock()
	for _, w := range p.poolMap {
		l = append(l, w)
	}

	return l
}

func (p *WorkPool) NewPool(name string, work int) *Pool {
	return p.NewPoolWithFunc(name, work, nil)
}

func (p *WorkPool) NewPoolWithFunc(name string, work int, logic Logic) *Pool {
	id := log.GenTraceId()

	pool := newPool(name, id, p)
	pool.SetWorker(work)
	pool.SetDefLogic(logic)

	p.poolLock.Lock()
	p.poolMap[id] = pool
	p.poolLock.Unlock()

	p.log.Infof("new pool %s(%s)", name, id)

	return pool
}

func (p *WorkPool) freePool(id string) {
	p.poolLock.Lock()
	delete(p.poolMap, id)
	p.poolLock.Unlock()
}

func (p *WorkPool) applyWorker() bool {
	p.workerLock.Lock()
	defer p.workerLock.Unlock()

	if p.worker.Load() >= p.maxWorker.Load() {
		return false
	}

	p.log.Infof("now has %d worker pool", p.worker.Inc())

	return true
}

func (p *WorkPool) applyWorkerForce() bool {
	p.workerLock.Lock()
	defer p.workerLock.Unlock()

	if p.worker.Load() >= p.maxWorker.Load() {
		return false
	}

	p.log.Infof("(force)now has %d worker pool", p.worker.Inc())

	return true
}

func (p *WorkPool) freeWorker() {
	p.worker.Dec()
}

func (p *WorkPool) LoadTotal() uint64 {
	var total uint64
	for _, w := range p.clonePool() {
		if w == nil {
			continue
		}
		p.log.Infof("%s(%s) worker:%d|max:%d|total:%d|wait:%d",
			w.name, w.id, w.worker.Load(), w.maxWorker.Load(), w.taskTotal.Load(), len(w.wait))
		total += w.taskTotal.Load()
	}

	return total
}

func (p *WorkPool) Close() {
	for _, w := range p.clonePool() {
		if w == nil {
			continue
		}
		p.log.Infof("close pool %s(%s)", w.name, w.id)
		w.Close()
	}

	close(p.subWorkerCloseChan)
}

func (p *WorkPool) onWorkerFree() {
	p.subWorkerCloseChan <- true
}

func (p *WorkPool) Statistics() Statistics {
	statistics := Statistics{
		WorkStatisticsMap: map[string]*WorkStatistics{},
	}

	for _, w := range p.clonePool() {
		if w == nil {
			continue
		}
		statistics.TotalWork += w.worker.Load()
		statistics.TotalTask += w.taskTotal.Load()
		statistics.TotalWait += uint64(len(w.wait))

		statistics.WorkStatisticsMap[fmt.Sprintf("%s(%s)", w.name, w.id)] = &WorkStatistics{
			Id:                w.id,
			Name:              w.name,
			TotalTask:         w.taskTotal.Load(),
			TotalWork:         w.worker.Load(),
			TotalWait:         uint64(len(w.wait)),
			AvgProcessingTime: time.Millisecond * time.Duration(w.avgTime.Value()),
		}
	}

	return statistics
}

var _pool *WorkPool

func init() {
	_pool = NewPool(32)
	_pool.SetLogger(log.Clone().SetOutput(ioutil.Discard))
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

func PoolGlobalStatistics() Statistics {
	return _pool.Statistics()
}

func DisablePoolGlobal() {
	_pool.Close()
}
