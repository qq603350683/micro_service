package handler

import (
	"context"
	"github.com/micro/go-micro/v2"
	"testing"
	webApiProto "webApi/proto/webApi"
)

func BenchmarkWebApi_GetInfoByUserId(b *testing.B) {
	b.StopTimer()

	service := micro.NewService(
		micro.Name("go.micro.api.webApi"),
		micro.Version("latest"),
		//micro.Registry(consul2),
	)

	service.Init()

	// Register Handler
	webApiProto.RegisterWebApiHandler(service.Server(), new(WebApi))

	// Run service
	if err := service.Run(); err != nil {
		b.Fatal(err)
	}

	web := &WebApi{}

	//num := rand.Intn(99)

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		get := make(map[string]*webApiProto.Pair)

		get["user_id"] = &webApiProto.Pair{
			Key:    "user_id",
			Values: []string{"1"},
		}

		req := &webApiProto.Request{
			Method: "GET",
			Path:   "",
			Header: nil,
			Get:    get,
			Post:   nil,
			Body:   "",
			Url:    "",
		}

		res := &webApiProto.Response{}

		err := web.GetInfoByUserId(context.TODO(), req, res)
		if err != nil {
			b.Error(err)
		}
	}
}