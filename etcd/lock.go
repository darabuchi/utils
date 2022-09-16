package etcd

import (
	"context"
	"errors"
	"time"

	"github.com/darabuchi/log"
	"go.etcd.io/etcd/client/v3/concurrency"
)

type Mutex struct {
	key     string
	timeout time.Duration
	lock    *concurrency.Mutex
}

func NewMutex(key string, timeout time.Duration) *Mutex {
	p := Mutex{
		key:     key,
		timeout: timeout,
	}

	return &p
}

func (p *Mutex) Lock() error {
	session, err := concurrency.NewSession(cli, concurrency.WithTTL(int(p.timeout.Seconds())))
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}

	p.lock = concurrency.NewMutex(session, p.key)

	err = p.lock.Lock(context.TODO())
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}

	return nil
}

func (p *Mutex) LockWithTimeout(timeout time.Duration) error {
	session, err := concurrency.NewSession(cli, concurrency.WithTTL(int(p.timeout.Seconds())))
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}

	p.lock = concurrency.NewMutex(session, p.key)

	timer := time.NewTimer(timeout)
	defer timer.Stop()
	for {
		err = p.lock.TryLock(context.TODO())
		if err != nil {
			if err != concurrency.ErrLocked {
				log.Errorf("err:%v", err)
				return err
			}

		} else {
			return nil
		}

		select {
		case <-time.After(time.Millisecond * 100):
			break
		case <-timer.C:
			return errors.New("lock timeout")
		}
	}
}

func (p *Mutex) Unlock() error {
	err := p.lock.Unlock(context.TODO())
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}

	return nil
}
