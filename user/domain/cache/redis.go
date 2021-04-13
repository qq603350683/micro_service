package cache

import (
	"github.com/go-redis/redis"
	"math/rand"
	"strconv"
	"time"
	"user/domain/config"
)

var RedisClient *redis.Client

func InitRedis(config *config.RedisConfig) {
	// go-redis 配置所有参数详细说明 https://blog.csdn.net/pengpengzhou/article/details/105385666
	var opts redis.Options

	opts.PoolSize = config.PoolSize
	opts.Addr = config.Addr
	opts.Password = config.Password
	opts.DB = config.DB

	RedisClient = redis.NewClient(&opts)
}


// 获取 Reids 锁
func Lock(key string) (bool, int) {
	randValue := rand.Intn(9999)

	b, err := RedisClient.SetNX(key, randValue, time.Duration(30 * time.Second)).Result()
	if err != nil && err != redis.Nil {
		return false, -1
	}

	return b, randValue
}


func UnLock(key string, randValue int) bool {
	// 释放锁
	result, err := RedisClient.Get(key).Result()
	if err != nil {
		return false
	}

	// 判断锁的值是否与传入的值一致，一致才能释放锁
	i, err := strconv.Atoi(result)
	if err != nil {
		return false
	}

	if i == randValue {
		i2, err := RedisClient.Del(key).Result()
		if err != nil {
			return false
		}

		if i2 <= 0 {
			return false
		}

		return true
	} else {
		return false
	}
}
