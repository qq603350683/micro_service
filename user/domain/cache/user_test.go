package cache

import (
	"testing"
	"time"
	"user/domain/model"
)

func TestUserCache_Add(t *testing.T) {
	userCache := NewUserCache()

	user := new(model.User)
	user.userId = 1
	user.UniqueID = "222"
	user.Nickname = "hello world"
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err := userCache.Add(user.userId, user)
	if err != nil {
		t.Error(err)
	}

	_, err = userCache.Add(2, nil)
	if err != nil {
		t.Error(err)
	}
}

func TestUserCache_Delete(t *testing.T) {
	userCache := NewUserCache()

	b, err := userCache.Delete(2)
	if err != nil {
		t.Error(err)
	}

	if b != true {
		t.Error("删除失败")
	}
}

func TestUserCache_GetInfoByUserId(t *testing.T) {
	userCache := NewUserCache()

	user, err := userCache.GetInfoByUserId(1)
	if err != nil {
		t.Error(err)
	}

	t.Log(user)
}

func TestUserCache_GetListByUserId(t *testing.T) {
	userCache := NewUserCache()

	list, err := userCache.GetListByUserId([]int{1, 2})
	if err != nil {
		t.Error(err)
	}

	t.Log(list)
}

func TestUserCache_AddUniqueID(t *testing.T) {
	userCache := NewUserCache()

	_, err := userCache.AddUniqueID("222", 1)
	if err != nil {
		t.Error(err)
	}
}

func TestUserCache_CheckUniqueIDIsExists(t *testing.T) {
	userCache := NewUserCache()

	b, err := userCache.CheckUniqueIDIsExists("222")
	if err != nil {
		t.Error(err)
	}

	if b != true {
		t.Error("测试与结果不一致")
	}

	b, err = userCache.CheckUniqueIDIsExists("333")
	if err != nil {
		t.Error(err)
	}

	if b == true {
		t.Error("测试与结果不一致")
	}
}

func TestUserCache_GetInfoByUniqueID(t *testing.T) {
	userCache := NewUserCache()

	user, err := userCache.GetInfoByUniqueID("222")
	if err != nil {
		t.Error(err)
	}

	t.Log(user)
}