module user

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require github.com/micro/go-micro/v2 v2.9.1

require github.com/jinzhu/gorm v1.9.12

require github.com/go-redis/redis v6.15.7+incompatible

require github.com/micro/go-plugins/config/source/consul/v2 v2.9.1 // indirect

require github.com/micro/go-plugins/registry/consul/v2 v2.9.1 // indirect

require github.com/micro/go-plugins/wrapper/trace/opentracing/v2 v2.9.1 // indirect

require github.com/opentracing/opentracing-go v1.2.0 // indirect

require github.com/uber/jaeger-client-go v2.25.0+incompatible // indirect

require (
	github.com/golang/protobuf v1.4.0
	github.com/micro/go-micro v1.18.0
	github.com/uber/jaeger-lib v2.4.0+incompatible // indirect
	// github.com/prometheus/common v0.6.0
	// github.com/joho/godotenv v1.3.0 // indirect
	// github.com/prometheus/common v0.6.0
	google.golang.org/protobuf v1.22.0
)
