package config

import (
	"github.com/micro/go-micro/v2/config"
)

//{
//"addr": "127.0.0.1:6379",
//"password": "",
//"db": 0,
//"pool_size": 5
//}

// Redis 配置
type RedisConfig struct {
	Addr string `json:"addr"`
	Password string `json:"password"`
	DB int `json:"db"`
	PoolSize int `json:"pool_size"`
}

func NewRedisConfig() *RedisConfig {
	return new(RedisConfig)
}

func GetRedisConfigFromConsul(config config.Config, path ...string) (*RedisConfig, error) {
	c := NewRedisConfig()

	err := config.Get(path...).Scan(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
