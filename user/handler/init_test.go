package handler

import (
	"testing"
	"user/domain/cache"
	"user/domain/repository"
)

func TestAll(t *testing.T) {

}

func TestMain(m *testing.M) {
 	repository.Database("root:root@(127.0.0.1:3306)/test_micro?charset=utf8&parseTime=True&loc=Local")

	cache.InitRedis()

	m.Run()
}