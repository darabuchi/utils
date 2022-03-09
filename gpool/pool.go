package gpool

import (
	"fmt"
	"github.com/VividCortex/ewma"
	"github.com/aofei/sandid"
	"github.com/darabuchi/log"
	"github.com/darabuchi/utils"
	"go.uber.org/atomic"
	"os"
	"sync"
	"time"
)

type Pool struct {
	id   string
	name string

	defLogic Logic

	pool     *workPool
	workLock sync.Mutex
	dataLock sync.RWMutex

	worker, maxWorker *atomic.Uint32
	wait              chan poolQueue

	taskTotal *atomic.Uint64

	stop chan bool

	taskWait, poolWait sync.WaitGroup

	healthAt time.Time
	timeout  time.Duration
	avgTime  ewma.MovingAverage
}

func newPool(name string, id string, pool *workPool) *Pool {
	return &Pool{
		id:        id,
		pool:      pool,
		maxWorker: atomic.NewUint32(3),
		worker:    atomic.NewUint32(0),
		name:      name,

		taskTotal: atomic.NewUint64(0),
		wait:      make(chan poolQueue, 100),
		stop:      make(chan bool, 1),

		healthAt: time.Now(),
		timeout:  time.Minute * 5,

		avgTime: ewma.NewMovingAverage(3),
	}
}

func (p *Pool) SetDefLogic(logic Logic) *Pool {
	p.dataLock.Lock()
	defer p.dataLock.Unlock()

	p.defLogic = logic
	return p
}

func (p *Pool) SetWorker(worker int) {
	if worker <= 0 {
		panic("worker master over zero")
	}
	p.workLock.Lock()
	p.maxWorker.Store(uint32(worker))
	p.workLock.Unlock()

	p.resetWorker()
}

func (p *Pool) resetWorker() {
	for {
		if !(p.maxWorker.Load() <= p.worker.Load()) {
			return
		}

		p.stop <- true
	}
}

func (p *Pool) run() {
	logic := func(i poolQueue) {
		start := time.Now()
		defer func() {
			p.avgTime.Add(float64(time.Since(start).Milliseconds()))
		}()
		defer func() {
			p.pool.onWorkerFree()
			p.taskTotal.Dec()
			p.taskWait.Done()
		}()
		defer utils.CachePanic()

		if i.logic == nil {
			return
		}

		if i.logic == nil {
			log.Error("logic is empty")
			return
		}

		i.logic(i.args)
	}

	p.poolWait.Add(1)
	go func(sign chan os.Signal) {
		defer p.poolWait.Done()

		traceId := fmt.Sprintf("%s.%s.%s", p.name, p.id, sandid.New().String())
		log.SetTrace(traceId)
		defer log.DelTrace()

		log.Infof("new worker start %s(%s)", p.name, p.id)
		defer func() {
			p.worker.Dec()
			p.pool.freeWorker()
			log.Warnf("%s(%s) worker is exist,work:%d|max:%d|wait:%d", p.name, p.id, p.worker.Load(), p.maxWorker.Load(), len(p.wait))
		}()
		defer utils.CachePanic()

		freeTicker := time.NewTimer(time.Second * 30)
		defer freeTicker.Stop()

		for {
			select {
			case i, ok := <-p.wait:
				if !ok {
					log.Warnf("%s(%s) pool is closed,free resource", p.name, p.id)
					return
				}

				log.SetTrace(fmt.Sprintf("%s.%s.%s", i.traceId, traceId, sandid.New().String()))
				logic(i)
				log.SetTrace(traceId)
				p.updateHealth()

				freeTicker.Reset(time.Second * 30)
			case <-time.After(time.Second * 3):
				if len(p.wait) == 0 {
					log.Warnf("%s(%s) without wait task,free resource", p.name, p.id)
					return
				}
				//log.Infof("wait timeout,free resource for %s(%s)", p.name, p.id)
				log.Infof("%s(%s) worker is still alive,work:%d|max:%d|wait:%d", p.name, p.id, p.worker.Load(), p.maxWorker.Load(), len(p.wait))
			case <-freeTicker.C:
				log.Warnf("wait timeout,free resource for %s(%s)", p.name, p.id)
				return
			case <-p.stop:
				log.Warnf("stop worker for %s(%s)", p.name, p.id)
				return
			case <-sign:
				log.Warnf("exist worker for %s(%s)", p.name, p.id)
				return
			}
		}
	}(utils.GetExitSign())
}

func (p *Pool) tryApply() bool {
	p.workLock.Lock()
	defer p.workLock.Unlock()

	if !p.needMoreWorkerWithoutLock() {
		return false
	}

	if !p.pool.applyWorker() {
		log.Infof("%s(%s) apply resource fail", p.name, p.id)
		// 申请资源失败，退出
		return false
	}

	// 申请成功了
	log.Infof("%s(%s) has %d worker pool", p.name, p.id, p.worker.Inc())

	p.run()
	return true
}

func (p *Pool) applyForce() {
	p.workLock.Lock()
	defer p.workLock.Unlock()

	// 已经没有任务了，没有必要申请
	if len(p.wait) == 0 {
		return
	}

	if !p.pool.applyWorkerForce() {
		log.Infof("%s(%s) apply resource fail", p.name, p.id)
		// 申请资源失败，退出
		return
	}

	// 申请成功了
	log.Infof("%s(%s) has %d worker pool", p.name, p.id, p.worker.Inc())

	p.run()
}

func (p *Pool) free() {
	select {
	case p.stop <- true:
		log.Warnf("%s(%s) notify stop", p.name, p.id)
	case <-time.After(time.Second * 5):
		log.Warnf("%s(%s) notify stop timeout", p.name, p.id)
	}
}

func (p *Pool) afterSubmit() {
	p.updateHealth()
	p.tryApply()
}

func (p *Pool) updateHealth() {
	p.dataLock.Lock()
	p.healthAt = time.Now()
	p.dataLock.Unlock()
}

func (p *Pool) needMoreWorker() bool {
	p.workLock.Lock()
	defer p.workLock.Unlock()

	return p.needMoreWorkerWithoutLock()
}

func (p *Pool) needMoreWorkerForce() bool {
	if len(p.wait) == 0 {
		return false
	}

	timeout := p.timeout
	if p.avgTime.Value() > 0 {
		timeout = time.Duration(p.avgTime.Value()*1.5) * time.Millisecond
	} else if timeout <= time.Minute {
		timeout = time.Minute
	}

	if timeout <= 0 {
		timeout = time.Second * 10
	}

	if p.worker.Load()/2 == p.maxWorker.Load() && time.Since(p.healthAt) >= timeout {
		return true
	}

	return false
}

func (p *Pool) needMoreFree() bool {
	if len(p.wait) == 0 {
		return false
	}

	timeout := p.timeout
	if p.avgTime.Value() > 0 {
		timeout = time.Duration(p.avgTime.Value()*1.5) * time.Millisecond
	} else if timeout <= time.Minute {
		timeout = time.Minute
	}

	if timeout <= 0 {
		timeout = time.Second * 10
	}

	if p.worker.Load()/2 == p.maxWorker.Load() && time.Since(p.healthAt) >= timeout {
		return true
	}

	return false
}

func (p *Pool) needMoreWorkerWithoutLock() bool {
	if p.worker.Load()+1 > p.maxWorker.Load() {
		log.Debugf("%s(%s) pool is full, skip", p.name, p.id)
		// 已经满了，退出
		return false
	}

	if len(p.wait) == 0 {
		return false
	}

	if uint32(len(p.wait)) < p.worker.Load() {
		log.Debugf("%s(%s) still has resource, skip", p.name, p.id)
		return false
	}

	return true
}

func (p *Pool) needClose() bool {
	p.dataLock.RLock()
	defer p.dataLock.RUnlock()

	if p.timeout < 0 {
		return false
	}

	if len(p.wait) > 0 {
		return false
	}

	if time.Since(p.healthAt) > p.timeout {
		return true
	}

	return false
}

func (p *Pool) SetAlways() *Pool {
	p.SetTimeout(-1)
	return p
}

func (p *Pool) SetTimeout(timeout time.Duration) *Pool {
	p.dataLock.Lock()
	p.timeout = timeout
	p.dataLock.Unlock()
	return p
}

func (p *Pool) Submit(i interface{}) {
	p.dataLock.RLock()
	logic := p.defLogic
	p.dataLock.RUnlock()

	p.SubmitWithFunc(i, logic)
}

func (p *Pool) SubmitWithTimeout(i interface{}, timeout time.Duration) {
	p.dataLock.RLock()
	logic := p.defLogic
	p.dataLock.RUnlock()

	p.SubmitWithFunc(i, func(i interface{}) {
		done := make(chan bool)

		go func() {
			defer func() {
				select {
				case done <- true:
				case <-time.After(time.Second):
				}
			}()
			if logic == nil {
				return
			}

			logic(i)
		}()

		select {
		case <-done:
		case <-time.After(timeout):
			log.Infof("timeout, skip")
		}
	})
}

func (p *Pool) SubmitWait(i interface{}) {
	p.dataLock.RLock()
	logic := p.defLogic
	p.dataLock.RUnlock()

	p.SubmitWithFuncWait(i, logic)
}

func (p *Pool) SubmitWithFunc(i interface{}, logic Logic) {
	log.Infof("submit task for %s(%s)", p.name, p.id)

	defer utils.CachePanic()

	p.taskTotal.Inc()
	p.taskWait.Add(1)
	p.wait <- poolQueue{
		logic:   logic,
		args:    i,
		traceId: log.GetTrace(),
	}
	p.afterSubmit()
}

func (p *Pool) SubmitWithFuncWait(i interface{}, logic Logic) {
	var w sync.WaitGroup
	w.Add(1)
	p.SubmitWithFunc(i, func(i interface{}) {
		defer w.Done()
		if logic == nil {
			return
		}
		logic(i)
	})
	w.Wait()
}

func (p *Pool) Wait() {
	p.taskWait.Wait()
}

func (p *Pool) Close() {
	log.Infof("close pool %s(%s)", p.name, p.id)
	close(p.wait)
	close(p.stop)

	p.poolWait.Wait()

	p.pool.freePool(p.id)
}

func (p *Pool) NewSubPool(name string) *SubPool {
	return p.NewSubPoolWithFunc(name, p.defLogic)
}

func (p *Pool) NewSubPoolWithFunc(name string, logic Logic) *SubPool {
	return newSubPoolWithFunc(name, p, logic)
}
