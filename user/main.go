package main

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	microHystrix "github.com/micro/go-plugins/wrapper/breaker/hystrix/v2"
	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	opentracing2 "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/opentracing/opentracing-go"
	"net"
	"net/http"
	"user/domain/cache"
	"user/domain/common"
	"user/domain/config"
	"user/domain/repository"
	"user/handler"
	userPro "user/proto/user"
)

func main() {
	const MicroServiceName = "go.micro.service.user"
	const QPS = 3000

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

	// 熔断器
	hystrixStreamHandler := hystrix.NewStreamHandler()
	hystrixStreamHandler.Start()
	// 启动端口
	go func() {
		err = http.ListenAndServe(net.JoinHostPort("0.0.0.0", "9096"), hystrixStreamHandler)
		if err != nil {
			log.Error(err)
		}
	}()

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
		//添加熔断
		micro.WrapClient(microHystrix.NewClientWrapper()),
		// 添加限流
		micro.WrapClient(ratelimit.NewClientWrapper(QPS)),
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
	common.Info("micro service:[" + MicroServiceName + "] success start")

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
