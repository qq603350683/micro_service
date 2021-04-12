package cache

import (
	"fmt"
	"testing"
)

func TestAll(t *testing.T) {
	//t.Run("TestUserCache_Add", TestUserCache_Add)
	//t.Run("TestUserCache_GetListByUserId", TestUserCache_GetListByUserId)
}

func TestMain(m *testing.M) {
	fmt.Println("begin")
	InitRedis()

	m.Run()
}
