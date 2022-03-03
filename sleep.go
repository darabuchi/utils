package utils

import (
	"math/rand"
	"time"
)

func RandSleep(min time.Duration, max time.Duration) {
	if max < min {
		return
	}

	if min == max {
		time.Sleep(min)
		return
	}

	time.Sleep(time.Duration(rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(int64(max-min))) + min)
}
