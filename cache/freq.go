package cache

import (
	"time"

	"github.com/darabuchi/log"
)

func FreqCheck(key string, limit int64, timeout time.Duration) (bool, error) {
	val, err := client.Incr(key)
	if err != nil {
		log.Errorf("err:%s", err)
		return false, err
	}

	if val > limit {
		log.Warnf("key:%s over limit:%d val:%d", key, limit, val)
		return false, nil
	}

	if val == 1 {
		for i := 0; i < 3; i++ {
			_, err = client.Expire(key, int64(timeout.Seconds()))
			if err != nil {
				log.Errorf("err:%v", err)
				continue
			}
			break
		}
	}

	return true, nil
}

func FreqTryLock(key string, limit int64, timeout time.Duration) (time.Duration, error) {
	val, err := Incr(key)
	if err != nil {
		log.Errorf("err:%s", err)
		return timeout / 2, err
	}

	if val > limit {
		log.Warnf("key:%s over limit:%d val:%d", key, limit, val)

		ttl, err := Ttl(key)
		if err != nil {
			log.Errorf("err:%v", err)
			return timeout, nil
		}

		if ttl < time.Second {
			ttl = time.Second
		}

		return ttl, nil
	}

	if val == 1 {
		for i := 0; i < 3; i++ {
			_, err = client.Expire(key, int64(timeout.Seconds()))
			if err != nil {
				log.Errorf("err:%v", err)
				continue
			}
			break
		}
		return 0, nil
	}

	return 0, nil
}
