package cache

import (
	"time"

	"github.com/darabuchi/log"
)

func TryLock(key string, timeout time.Duration) (bool, error) {
	return SetNxWithTimeout(key, 0, timeout)
}

func Unlock(key string) error {
	_, err := Del(key)
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}
	return nil
}

func Lock(key string, timeout time.Duration) {
	for {
		ok, err := TryLock(key, timeout)
		if err != nil {
			log.Errorf("err:%v", err)
			time.Sleep(time.Second)
			continue
		}

		if ok {
			break
		}

		wait, err := Ttl(key)
		if err != nil {
			log.Errorf("err:%v", err)
			time.Sleep(time.Second)
			continue
		}

		time.Sleep(wait)
	}
}
