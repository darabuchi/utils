package utils

import (
	"math/rand"
	"time"
)

func RandSleep(min time.Duration, max time.Duration) {
	time.Sleep(RandDuration(min, max))
}

func RandDuration(min time.Duration, max time.Duration) time.Duration {
	if max < min {
		return min
	}

	if min == max {
		return min
	}

	return time.Duration(rand.New(rand.NewSource(time.Now().UnixNano())).Int63n(int64(max-min))) + min
}
