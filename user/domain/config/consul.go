package config

import (
	"fmt"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-plugins/config/source/consul/v2"
)

type ConsulConfig struct {
	IP string
	Port int64
	Addr string
}

func NewConsulConfig() *ConsulConfig {
	c := new(ConsulConfig)

	c.IP = "192.168.18.128"
	c.Port = 8500
	c.Addr = fmt.Sprintf("%s:%d", c.IP, c.Port)

	return c
}

func GetConsulConfig(host string, port int64, prefix string)  (config.Config, error) {
	address := fmt.Sprintf("%s:%d", host, port)

	consulSource := consul.NewSource(
		// 设置配置中心地址
		consul.WithAddress(address),
		// 设置前缀，不设置默认前缀为  /micro/config
		consul.WithPrefix(prefix),
		// 是否移除前缀，这里设置为 true， 表示可以不带前缀直接获取对应配置
		consul.StripPrefix(true),
	)

	// 配置初始化
	config, err := config.NewConfig()
	if err != nil {
		return config, err
	}

	// 加载配置
	err = config.Load(consulSource)

	return config, err
}
