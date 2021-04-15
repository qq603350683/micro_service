package main

import (
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"webApi/handler"
	webApi "webApi/proto/webApi"
)

func main() {
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

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
