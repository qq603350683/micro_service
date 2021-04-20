module webApi

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require github.com/micro/go-micro/v2 v2.9.1

require github.com/micro/go-plugins/config/source/consul/v2 v2.9.1

require github.com/micro/go-plugins/registry/consul/v2 v2.9.1

require github.com/micro/go-plugins/wrapper/trace/opentracing/v2 v2.9.1

require github.com/opentracing/opentracing-go v1.2.0

require github.com/uber/jaeger-client-go v2.25.0+incompatible

require (
	github.com/golang/protobuf v1.4.0
	github.com/prometheus/common v0.6.0
	github.com/qq603350683/micro_service v0.0.0-20210415094550-411582d251ec // indirect
	github.com/uber/jaeger-lib v2.4.0+incompatible // indirect
	go.uber.org/zap v1.16.0
	google.golang.org/genproto v0.0.0-20191216164720-4f79533eabd1
	// github.com/prometheus/common v0.6.0
	// github.com/joho/godotenv v1.3.0 // indirect
	// github.com/prometheus/common v0.6.0
	google.golang.org/protobuf v1.22.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)
