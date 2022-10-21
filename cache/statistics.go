package cache

import (
	"fmt"
	"time"

	"github.com/darabuchi/log"
	"github.com/darabuchi/utils/xtime"
)

func Statistics(key string, duration time.Duration) {
	switch duration {
	case xtime.Day:
		key = fmt.Sprintf("statistics:%s:day:%s", key, time.Now().Format("20060102"))
	}

	cnt, err := Incr(key)
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}

	if cnt == 1 {
		switch duration {
		case xtime.Day:
			_, err = Expire(key, xtime.Week)
			if err != nil {
				log.Errorf("err:%v", err)
			}
		}
	}
}
