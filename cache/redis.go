package cache

import (
	"strconv"
	"time"

	"github.com/bytedance/sonic"
	"github.com/darabuchi/log"
	"github.com/darabuchi/utils"
	"github.com/garyburd/redigo/redis"
	"github.com/shomali11/xredis"
)

var client *xredis.Client

func Connect(addr string, db int, password string) error {
	if client != nil {
		return nil
	}

	client = xredis.NewClient(&redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", addr,
				redis.DialDatabase(1),
				redis.DialConnectTimeout(time.Second*3),
				redis.DialReadTimeout(time.Second*3),
				redis.DialWriteTimeout(time.Second*3),
				redis.DialKeepAlive(time.Minute),
				redis.DialPassword(password),
			)
		},
		MaxIdle:     100,
		MaxActive:   100,
		IdleTimeout: time.Second * 5,
		Wait:        true,
	})

	ping, err := client.Ping()
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}

	log.Infof("ping %s", ping)

	return nil
}

func Incr(key string) (int64, error) {
	return client.Incr(key)
}

func Decr(key string) (int64, error) {
	return client.Decr(key)
}

func IncrBy(key string, value int64) (int64, error) {
	return client.IncrBy(key, value)
}

func IncrByFloat(key string, increment float64) (float64, error) {
	return client.IncrByFloat(key, increment)
}

func DecrBy(key string, value int64) (int64, error) {
	return client.DecrBy(key, value)
}

func Get(key string) (string, error) {
	val, ok, err := client.Get(key)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", redis.ErrNil
	}

	return val, nil
}

func GetUint64(key string) (uint64, error) {
	val, err := Get(key)
	if err != nil {
		if err == redis.ErrNil {
			return 0, nil
		}

		return 0, err
	}

	num, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return 0, err
	}

	return num, nil
}

func Exists(keys ...string) (bool, error) {
	ok, err := client.Exists(keys...)
	if err != nil {
		if err == redis.ErrNil {
			return false, nil
		}

		return false, err
	}

	return ok, nil
}

func SetNx(key string, value interface{}) (bool, error) {
	ok, err := client.SetNx(key, utils.ToString(value))
	if err != nil {
		if err == redis.ErrNil {
			return false, nil
		}

		return false, err
	}

	return ok, nil
}

func Expire(key string, timeout time.Duration) (bool, error) {
	ok, err := client.Expire(key, int64(timeout.Seconds()))
	if err != nil {
		if err == redis.ErrNil {
			return false, nil
		}

		return false, err
	}

	return ok, nil
}

func GetJson(key string, j interface{}) error {
	val, err := Get(key)
	if err != nil {
		if err != redis.ErrNil {
			log.Errorf("err:%s", err)
		}
		return err
	}
	err = sonic.Unmarshal([]byte(val), j)
	if err != nil {
		log.Errorf("err:%s", err)
		return err
	}

	return nil
}

func GetInt64(key string) (int64, error) {
	val, err := Get(key)
	if err != nil {
		if err == redis.ErrNil {
			return 0, nil
		}

		return 0, err
	}

	num, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, err
	}

	return num, nil
}

func GetFloat64(key string) (float64, error) {
	val, err := Get(key)
	if err != nil {
		if err == redis.ErrNil {
			return 0, nil
		}

		return 0, err
	}

	num, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return 0, err
	}

	return num, nil
}

func Ttl(key string) (time.Duration, error) {
	connection := client.GetConnection()
	defer connection.Close()

	ttl, err := redis.Int64(connection.Do("TTL", key))
	if err != nil {
		log.Errorf("err:%v", err)
		return 0, err
	}

	return time.Duration(ttl) * time.Second, nil
}

func Set(key string, value interface{}) (bool, error) {
	return client.Set(key, utils.ToString(value))
}

func SetEx(key string, value interface{}, timeout time.Duration) (bool, error) {
	return client.SetEx(key, utils.ToString(value), int64(timeout.Seconds()))
}

func SetNxWithTimeout(key string, value interface{}, timeout time.Duration) (bool, error) {
	ok, err := SetNx(key, value)
	if err != nil {
		log.Errorf("err:%v", err)
		return false, err
	}

	if ok {
		_, err = Expire(key, timeout)
		if err != nil {
			log.Errorf("err:%v", err)
		}
	}

	return ok, nil
}

func Del(keys ...string) (int64, error) {
	return client.Del(keys...)
}

func HSet(key string, field string, value interface{}) (bool, error) {
	return client.HSet(key, field, utils.ToString(value))
}

func HGetAll(key string) (map[string]string, error) {
	return client.HGetAll(key)
}

func HKeys(key string) ([]string, error) {
	return client.HKeys(key)
}

func HVals(key string) ([]string, error) {
	conn := client.GetConnection()
	defer conn.Close()

	return redis.Strings(conn.Do("HVALS", key))
}

func HDel(key string, fields ...string) (int64, error) {
	return client.HDel(key, fields...)
}

func SAdd(key string, members ...string) (int64, error) {
	conn := client.GetConnection()
	defer conn.Close()

	args := make([]interface{}, 0, len(members)+1)
	args = append(args, key)
	for _, member := range members {
		args = append(args, member)
	}

	return redis.Int64(conn.Do("SADD", args...))
}

func SMembers(key string) ([]string, error) {
	conn := client.GetConnection()
	defer conn.Close()

	return redis.Strings(conn.Do("SMEMBERS", key))
}

func SRem(key string, members ...string) (int64, error) {
	conn := client.GetConnection()
	defer conn.Close()

	args := make([]interface{}, 0, len(members)+1)
	args = append(args, key)
	for _, member := range members {
		args = append(args, member)
	}

	return redis.Int64(conn.Do("SREM", args...))
}

func HIncr(key string, subKey string) (int64, error) {
	return client.HIncr(key, subKey)
}

func HIncrBy(key string, field string, increment int64) (int64, error) {
	return client.HIncrBy(key, field, increment)
}

func HIncrByFloat(key string, field string, increment float64) (float64, error) {
	return client.HIncrByFloat(key, field, increment)
}

func HDecr(key string, field string) (int64, error) {
	return client.HDecr(key, field)
}

func HDecrBy(key string, field string, increment int64) (int64, error) {
	return client.HDecrBy(key, field, increment)
}
