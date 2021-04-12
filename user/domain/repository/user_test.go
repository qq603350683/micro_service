package repository

import (
	"fmt"
	"testing"
	"user/domain/model"
)

func TestPipeline(t *testing.T) {
	var err error

	U := NewUserRepository()

	user := model.NewUser()
	user.UniqueID = "1"
	user.Nickname = "2"
	user.Avatar = "https://www.avatar.com/333.jpg"

	err = U.Add(user)
	if err != nil {
		t.Error(err)
	}

	user.UniqueID = "2"
	user.Nickname = "3"

	err = U.Update(user)
	if err != nil {
		t.Error(err)
	}

	user1, err := U.GetInfoByUniqueID(user.UniqueID, "")
	if err != nil {
		t.Error(err)
	}

	if user1 == nil {
		t.Error("user1 is nil")
	}

	user2, err := U.GetInfoByUserId(user.userId, "")
	if err != nil {
		t.Error(err)
	}

	if user2 == nil {
		t.Error("user2 is nil")
	}

	users, err := U.GetListByUserId([]int{user.userId}, "")
	if err != nil {
		t.Error(err)
	}

	if len(users) == 0 {
		t.Error("users len is 0")
	}

	err = U.Delete(user.userId)
	if err != nil {
		t.Error(err)
	}
}

func TestUserRepository_Add(t *testing.T) {
	U := NewUserRepository()

	user := model.NewUser()
	user.UniqueID = "epoasvionasopwqfnimaod"
	user.Nickname = "hello"
	user.Avatar = "https://www.avatar.com/1.jpg"

	err := U.Add(user)
	if err != nil {
		t.Error(err)
	}
}

func TestUserRepository_Update(t *testing.T) {
	U := NewUserRepository()

	user := model.NewUser()
	user.userId = 1
	user.UniqueID = "xxxxxxxxxxxxxxx"
	user.Nickname = "bbbb"
	user.Avatar = "https://www.avatar.com/999.jpg"

	err := U.Update(user)
	if err != nil {
		t.Error(err)
	}
}

func TestUserRepository_Delete(t *testing.T) {
	U := NewUserRepository()

	err := U.Delete(1)
	if err != nil {
		t.Error(err)
	}
}

func TestUserRepository_GetInfoByOpenid(t *testing.T) {
	U := NewUserRepository()

	user, err := U.GetInfoByUniqueID("xxxxxxxxxxxxxxx", "")
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%v", user)

	user, err = U.GetInfoByUniqueID("123123123", "")
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%v", user)
}

func TestUserRepository_GetInfoByUserId(t *testing.T) {
	U := NewUserRepository()

	user, err := U.GetInfoByUserId(1, "user_id")
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%v", user)
}

func TestUserRepository_GetListByUserId(t *testing.T) {
	U := NewUserRepository()

	users, err := U.GetListByUserId([]int{1, 2}, "*")
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%v", users)
}