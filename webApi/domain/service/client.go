package service

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"time"
)

const ConsulAddr = "192.168.18.128:8500"

func GetClient() client.Client {
	// 注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			ConsulAddr,
		}
	})

	s := micro.NewService()

	s.Init(micro.Registry(consulRegistry))

	cli := s.Client()

	// 这里很奇怪，好像没啥作用 time.Second * 30
	opts := cli.Options()
	opts.CallOptions.RequestTimeout = time.Second * 30
	opts.CallOptions.DialTimeout = time.Second * 30

	return cli
}