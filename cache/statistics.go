package cache

import (
	"fmt"
	"time"

	"github.com/darabuchi/log"
	"github.com/darabuchi/utils/xtime"
)

func Statistics(key string, duration time.Duration) {
	switch duration {
	case xtime.Hour:
		key = fmt.Sprintf("statistics:%s:hour:%s", key, time.Now().Format("2006010215"))
	case xtime.Day:
		key = fmt.Sprintf("statistics:%s:day:%s", key, time.Now().Format("20060102"))
	// case xtime.Week:
	// 	key = fmt.Sprintf("statistics:%s:weeb:%s", key, time.Now().Format("20060102"))
	case xtime.Month:
		key = fmt.Sprintf("statistics:%s:month:%s", key, time.Now().Format("200601"))
	}

	cnt, err := Incr(key)
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}

	if cnt == 1 {
		switch duration {
		case xtime.Hour:
			_, err = Expire(key, xtime.Day*3)
			if err != nil {
				log.Errorf("err:%v", err)
			}
		case xtime.Day:
			_, err = Expire(key, xtime.Week)
			if err != nil {
				log.Errorf("err:%v", err)
			}
		case xtime.Month:
			_, err = Expire(key, xtime.Month+xtime.Week)
			if err != nil {
				log.Errorf("err:%v", err)
			}
		}
	}
}
