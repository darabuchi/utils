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
		return fmt.Sprintf("statistics:%s:hour:%s", key, t.Format("2006010215"))
	case xtime.Day:
		return fmt.Sprintf("statistics:%s:day:%s", key, t.Format("20060102"))
	case xtime.Week:
		year, week := t.ISOWeek()
		return fmt.Sprintf("statistics:%s:week:%04d%02d", key, year, week)
	case xtime.Month:
		return fmt.Sprintf("statistics:%s:month:%s", key, t.Format("200601"))
	case xtime.Year:
		return fmt.Sprintf("statistics:%s:year:%s", key, t.Format("2006"))
	default:
		return key
	}
}

func StatisticsIncr(key string, durations ...time.Duration) {
	for _, duration := range durations {
		statisticsIncrBy(key, duration, 1)
	}
}

func StatisticsIncrBy(key string, val int64, durations ...time.Duration) {
	for _, duration := range durations {
		statisticsIncrBy(key, duration, val)
	}
}

func StatisticsIncrByFloat(key string, val float64, durations ...time.Duration) {
	for _, duration := range durations {
		statisticsIncrByFloat(key, duration, val)
	}
}

func statisticsIncrBy(key string, duration time.Duration, val int64) {
	key = getStatisticsKey(key, duration, time.Now())

	cnt, err := IncrBy(key, val)
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}

	if cnt == val {
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
		case xtime.Week:
			_, err = Expire(key, xtime.Week+xtime.Day*3)
			if err != nil {
				log.Errorf("err:%v", err)
			}
		case xtime.Month:
			_, err = Expire(key, xtime.Month+xtime.Week)
			if err != nil {
				log.Errorf("err:%v", err)
			}
		case xtime.Year:
			_, err = Expire(key, xtime.Year+xtime.Week)
			if err != nil {
				log.Errorf("err:%v", err)
			}
		default:
			_, err = Expire(key, duration+xtime.Month)
			if err != nil {
				log.Errorf("err:%v", err)
			}
		}
	}
}

func statisticsIncrByFloat(key string, duration time.Duration, val float64) {
	key = getStatisticsKey(key, duration, time.Now())

	cnt, err := IncrByFloat(key, val)
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}

	if cnt == val {
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
		case xtime.Week:
			_, err = Expire(key, xtime.Week+xtime.Day*3)
			if err != nil {
				log.Errorf("err:%v", err)
			}
		case xtime.Month:
			_, err = Expire(key, xtime.Month+xtime.Week)
			if err != nil {
				log.Errorf("err:%v", err)
			}
		case xtime.Year:
			_, err = Expire(key, xtime.Year+xtime.Week)
			if err != nil {
				log.Errorf("err:%v", err)
			}
		default:
			_, err = Expire(key, duration+xtime.Month)
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

func GetStatisticsForFloat(key string, duration time.Duration, t time.Time) float64 {
	key = getStatisticsKey(key, duration, t)

	cnt, err := GetFloat64(key)
	if err != nil {
		log.Errorf("err:%v", err)
		return 0
	}
	return cnt
}
