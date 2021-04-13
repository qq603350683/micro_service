package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"user/handler"
	userPro "user/proto/user"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go_micro_service_user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	user := handler.NewUser()

	// Register Handler
	userPro.RegisterUserHandler(service.Server(), user)

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.user", service.Server(), new(subscriber.User))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("go.micro.srv.user", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
