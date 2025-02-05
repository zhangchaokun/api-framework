package gredis

import (
	"api-framework/core/config"
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	"time"
)

var RedisConn *redis.Pool

/**
 * =================================================================================================
 * Redis连接池
 * =================================================================================================
 */
func InitRedis() error {
	RedisConn = &redis.Pool{
		MaxIdle:     config.RedisConfig.MaxIdle,
		MaxActive:   config.RedisConfig.MaxActive,
		IdleTimeout: config.RedisConfig.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.RedisConfig.Host)
			if err != nil {
				return nil, err
			}
			if config.RedisConfig.Password != "" {
				if _, err := c.Do("AUTH", config.RedisConfig.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return nil
}

/**
 * =================================================================================================
 * String操作
 * =================================================================================================
 */
func Set(key string, data interface{}, time int) error {
	conn := RedisConn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}
	return nil
}

func Exists(key string) bool {
	conn := RedisConn.Get()
	defer conn.Close()
	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}
	return exists
}

func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	return redis.Bool(conn.Do("DEL", key))
}

func LikeDeletes(key string) error {
	conn := RedisConn.Get()
	defer conn.Close()
	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}
	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}
	return nil
}
