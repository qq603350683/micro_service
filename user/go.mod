module user

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require github.com/micro/go-micro/v2 v2.9.1

require github.com/jinzhu/gorm v1.9.12

require github.com/go-redis/redis v6.15.7+incompatible

require github.com/micro/go-plugins/config/source/consul/v2 v2.9.1

require github.com/micro/go-plugins/registry/consul/v2 v2.9.1

require github.com/micro/go-plugins/wrapper/trace/opentracing/v2 v2.9.1

require github.com/opentracing/opentracing-go v1.2.0

require github.com/uber/jaeger-client-go v2.25.0+incompatible

require (
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Masterminds/sprig v2.22.0+incompatible // indirect
	github.com/aokoli/goutils v1.1.1 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.5.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2
	github.com/google/uuid v1.2.0 // indirect
	github.com/huandu/xstrings v1.3.2 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/mitchellh/copystructure v1.1.2 // indirect
	github.com/mwitkow/go-proto-validators v0.3.2 // indirect
	github.com/pseudomuto/protoc-gen-doc v1.4.1 // indirect
	github.com/uber/jaeger-lib v2.4.0+incompatible // indirect
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2 // indirect
	google.golang.org/genproto v0.0.0-20210406143921-e86de6bf7a46 // indirect
	// github.com/prometheus/common v0.6.0
	// github.com/joho/godotenv v1.3.0 // indirect
	// github.com/prometheus/common v0.6.0
	google.golang.org/protobuf v1.26.0
)
