//go:build go1.18 || go1.19 || go1.20 || 1.18 || 1.19 || 1.20
// +build go1.18 go1.19 go1.20 1.18 1.19 1.20

package utils

import "sync"

func Async[M any](process int, push func(chan M), logic func(M)) {
	c := make(chan M)

	var w sync.WaitGroup
	w.Add(process)
	for i := 0; i < process; i++ {
		go func() {
			defer w.Done()

			var x M
			for x = range c {
				logic(x)
			}
		}()
	}

	push(c)
	close(c)

	w.Wait()
}
