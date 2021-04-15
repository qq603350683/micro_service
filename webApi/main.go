package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"webApi/handler"
	webApi "webApi/proto/webApi"
)

func main() {
	// 注册中心
	//consul2 := consul.NewRegistry(func(options *registry.Options) {
	//	options.Addrs = []string{
	//		"192.168.18.128:8500",
	//	}
	//})

	// New Service
	service := micro.NewService(
		micro.Name("go.micro.api.webApi"),
		micro.Version("latest"),
		//micro.Registry(consul2),
	)

	// Initialise service
	service.Init()

	// Register Handler
	webApi.RegisterWebApiHandler(service.Server(), new(handler.WebApi))

	handler.Service = service

	//req := service.Client().NewRequest("go.micro.service.lottery", "Lottery.AddLottery", map[string]string {
	//	"name": "123123",
	//	"begin_at": "2021-03-08",
	//	"end_at": "2021-03-18",
	//})
	//
	//var res interface{}
	//
	//service.Client().Call(context.TODO(), req, res)

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.service.webApi", service.Server(), new(subscriber.WebApi))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
