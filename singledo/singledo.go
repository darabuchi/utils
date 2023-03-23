package singledo

import (
	"sync"
	"time"
)

type call[M any] struct {
	wg  sync.WaitGroup
	val M
	err error
}

type Single[M any] struct {
	mux    sync.Mutex
	last   time.Time
	wait   time.Duration
	call   *call[M]
	result *Result[M]
}

type Result[M any] struct {
	Val M
	Err error
}

func (s *Single[M]) Do(fn func() (M, error)) (v M, shared bool, err error) {
	s.mux.Lock()
	now := time.Now()
	if now.Before(s.last.Add(s.wait)) {
		s.mux.Unlock()
		return s.result.Val, true, s.result.Err
	}

	if call := s.call; call != nil {
		s.mux.Unlock()
		call.wg.Wait()
		return call.val, true, call.err
	}

	call := &call[M]{}
	call.wg.Add(1)
	s.call = call
	s.mux.Unlock()
	call.val, call.err = fn()
	call.wg.Done()

	s.mux.Lock()
	s.call = nil
	s.result = &Result[M]{Val: call.val, Err: call.err}
	s.last = now
	s.mux.Unlock()
	return call.val, false, call.err
}

func (s *Single[M]) DoWithoutCacheOnFail(fn func() (M, error)) (v M, shared bool, err error) {
	s.mux.Lock()
	now := time.Now()
	if now.Before(s.last.Add(s.wait)) {
		s.mux.Unlock()
		return s.result.Val, true, s.result.Err
	}

	if call := s.call; call != nil {
		s.mux.Unlock()
		call.wg.Wait()
		return call.val, true, call.err
	}

	call := &call[M]{}
	call.wg.Add(1)
	s.call = call
	s.mux.Unlock()
	call.val, call.err = fn()
	call.wg.Done()

	if call.err != nil {
		return call.val, false, call.err
	}

	s.mux.Lock()
	s.call = nil
	s.result = &Result[M]{Val: call.val, Err: call.err}
	s.last = now
	s.mux.Unlock()
	return call.val, false, call.err
}

func (s *Single[M]) MustDoWithoutCacheOnFail(fn func() (M, error)) (v M, shared bool, err error) {
	s.mux.Lock()
	now := time.Now()

	call := &call[M]{}
	call.wg.Add(1)
	s.call = call
	s.mux.Unlock()
	call.val, call.err = fn()
	call.wg.Done()

	if call.err != nil {
		return call.val, false, call.err
	}

	s.mux.Lock()
	s.call = nil
	s.result = &Result[M]{Val: call.val, Err: call.err}
	s.last = now
	s.mux.Unlock()
	return call.val, false, call.err
}

func (s *Single[M]) DoFast(fn func() (M, error)) (v M, err error) {
	v, _, err = s.Do(fn)
	return
}

func (s *Single[M]) Reset() {
	s.mux.Lock()
	s.last = time.Time{}
	s.mux.Unlock()
}

func NewSingle[M any](wait time.Duration) *Single[M] {
	return &Single[M]{wait: wait}
}
