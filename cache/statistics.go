package cache

import (
	"fmt"
	"time"

	"github.com/darabuchi/log"
	"github.com/darabuchi/utils/xtime"
)

func Statistics(key string, durations ...time.Duration) {
	StatisticsIncr(key, durations...)
}

func getStatisticsKey(key string, duration time.Duration, t time.Time) string {
	switch duration {
	case xtime.Hour:
		return fmt.Sprintf("statistics:%s:hour:%s", key, time.Now().Format("2006010215"))
	case xtime.Day:
		return fmt.Sprintf("statistics:%s:day:%s", key, time.Now().Format("20060102"))
	case xtime.Week:
		year, week := t.ISOWeek()
		return fmt.Sprintf("statistics:%s:week:%04d%02d", key, year, week)
	case xtime.Month:
		return fmt.Sprintf("statistics:%s:month:%s", key, time.Now().Format("200601"))
	case xtime.Year:
		return fmt.Sprintf("statistics:%s:year:%s", key, time.Now().Format("2006"))
	default:
		return key
	}
}

func StatisticsIncr(key string, durations ...time.Duration) {
	for _, duration := range durations {
		statisticsIncrBy(key, duration)
	}
}

func statisticsIncrBy(key string, duration time.Duration) {
	key = getStatisticsKey(key, duration, time.Now())

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
		default:
			_, err = Expire(key, duration+xtime.Week)
			if err != nil {
				log.Errorf("err:%v", err)
			}
		}
	}
}

func GetStatistics(key string, duration time.Duration, t time.Time) int64 {
	key = getStatisticsKey(key, duration, t)

	cnt, err := GetInt64(key)
	if err != nil {
		log.Errorf("err:%v", err)
		return 0
	}
	return cnt
}
