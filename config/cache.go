package config

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/pelletier/go-toml"
	"time"
)

type RedisConfig struct {
	Host        string        `yaml:"host"`
	Port        string        `yaml:"port"`
	Pass        string        `yaml:"pass"`
	Database    int           `yaml:"database"`
	MaxIdle     int           `yaml:"maxIdle"`   // 最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除
	MaxActive   int           `yaml:"maxActive"` // 最大的激活连接数，表示同时最多有N个连接 ，为0事表示没有限制
	IdleTimeout time.Duration `yaml:"timeout"`   // 最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
}

var redisConnPool *redis.Pool

func DefaultRedisConfig() RedisConfig {
	return RedisConfig{Host: "127.0.0.1", Port: "6379"}
}

func InitRedis(configToml *toml.Tree) {
	redisConfig := DefaultRedisConfig()
	if configToml != nil {
		if err := configToml.Unmarshal(&redisConfig); err != nil {
			logger.Fatalf("Init redis err:{}", err)
		}
	}
	redisConnPool = InitRedisConnPool(&redisConfig)
}

func CacheSet(key string, val interface{}, ttl int) error {
	_, err := GetRedisClient().Do("SET", key, val, "EX", ttl)
	return err
}

func CacheSetJson(key string, val interface{}, ttl int) error {
	byts, err := json.Marshal(&val)
	if err == nil {
		_, err = GetRedisClient().Do("SET", key, byts, "EX", ttl)
	}
	return err
}

func CacheHset(key, field string, val interface{}, ttl int) error {
	byts, err := json.Marshal(&val)
	if err == nil {
		_, err = GetRedisClient().Do("hset", key, field, byts)
	}
	if err == nil {
		err = CacheExpire(key, ttl)
	}
	return err
}

func CacheGetbytes(key string) []byte {
	val, _ := redis.Bytes(GetRedisClient().Do("GET", key))
	return val
}

/**
 * struct2 注意大小写，需传入指针
 */
func CacheGetStruct(key string, struct2 interface{}) {
	var err error
	byts := CacheGetbytes(key)
	if byts != nil {
		err = json.Unmarshal(byts, struct2)
	}
	if err != nil {
		logger.Warn("cache: get struct key : {}, err:{}", key, err)
	}
}

func CacheHGetStruct(key, field string, struct2 interface{}) {
	var err error
	byts, _ := redis.Bytes(GetRedisClient().Do("hget", key, field))
	if byts != nil {
		err = json.Unmarshal(byts, struct2)
	}
	if err != nil {
		logger.Warn("cache: hget struct key : {}, err:{}", key, err)
	}
}

func CacheGetString(key string) string {
	val, err := redis.String(GetRedisClient().Do("GET", key))
	if err != nil {
		logger.Warnf("cache err when get key : {},err: {}", key, err)
		return ""
	}
	return val
}

func CacheExpire(key string, ttl int) error {
	_, err := GetRedisClient().Do("EXPIRE", key, ttl)
	if err != nil {
		logger.Warnf("cache err when expire key :", key)
	}
	return err
}

func CacheDelete(key string) error {
	_, err := GetRedisClient().Do("DEL", key)
	if err != nil {
		logger.Warnf("cache err when get key :", key)
	}
	return err
}

func GetRedisClient() redis.Conn {
	return redisConnPool.Get()
}

func InitRedisConnPool(redisConfig *RedisConfig) *redis.Pool {
	server := fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port)
	return &redis.Pool{
		MaxIdle:     redisConfig.MaxIdle,
		MaxActive:   redisConfig.MaxActive,
		IdleTimeout: redisConfig.IdleTimeout * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if redisConfig.Pass != "" {
				if _, err := c.Do("AUTH", redisConfig.Pass); err != nil {
					c.Close()
					return nil, err
				}
			}
			if _, err := c.Do("SELECT", redisConfig.Database); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}
