package main

import (
	"fmt"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"user/domain/cache"
	"user/domain/config"
	"user/domain/repository"
	"user/handler"
	userPro "user/proto/user"
)

var MicroServiceName string

func main() {
	MicroServiceName = "go.micro.service.user"

	consulConfigInfo := config.NewConsulConfig()
	jaegerConfigInfo := config.NewJaegerConfig()

	// 获取配置中心
	consulConfig, err := config.GetConsulConfig(consulConfigInfo.IP, consulConfigInfo.Port, "/micro/config/setting")
	if err != nil {
		log.Fatal(err)
	}

	// 获取 MySQL 配置信息并初始化
	mysql, err := config.GetMysqlConfigFromConsul(consulConfig, "mysql")
	if err != nil {
		log.Fatal(err)
	}
	conn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", mysql.User, mysql.Password, mysql.Host, mysql.Port, mysql.Database)
	repository.Database(conn)

	// 获取 Redis 配置并初始化
	redis, err := config.GetRedisConfigFromConsul(consulConfig, "redis")
	if err != nil {
		log.Fatal(err)
	}
	cache.InitRedis(redis)

	// 注册中心
	consulRegistry :=  consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			consulConfigInfo.Addr,
		}
	})

	// 链路追踪
	t, io, err := config.NewTracer(MicroServiceName, jaegerConfigInfo.Addr)
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// New Service
	service := micro.NewService(
		micro.Name(MicroServiceName),
		micro.Version("latest"),
		// 这里设置地址和需要暴露的端口
		micro.Address("127.0.0.1:8082"),
		// 这里添加 consul 作为注册中心
		micro.Registry(consulRegistry),
		// 绑定链路追踪 opentracing2
		micro.WrapHandler(opentracing2.NewHandlerWrapper(opentracing.GlobalTracer())),
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
