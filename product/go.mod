module product

go 1.13

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require github.com/micro/go-micro/v2 v2.9.1

require github.com/jinzhu/gorm v1.9.12

require github.com/go-redis/redis v6.15.7+incompatible

require github.com/micro/go-micro v1.18.0
