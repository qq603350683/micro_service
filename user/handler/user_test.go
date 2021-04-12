package handler

import (
	"context"
	"testing"
	userPro "user/proto/user"
)

func TestUser_Add(t *testing.T) {
	u := NewUser()

	req := &userPro.AddRequest{
		UniqueId: "22",
		Nickname: "hello",
		Avatar:   "1.jpg",
	}

	res := &userPro.AddResponse{}

	err := u.Add(context.TODO(), req, res)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(res)
	}
}

func TestUser_GetInfoByUserID(t *testing.T) {
	u := NewUser()

	req := &userPro.GetInfoByUserIDRequest{UserId:1}

	res := &userPro.GetInfoByUserIDResponse{}

	err := u.GetInfoByUserID(context.TODO(), req, res)
	if err != nil {
		t.Error(err)
	}

	t.Log(res)
}
