package handler

import (
	"context"
	"testing"
	userPro "user/proto/user"
)

func TestUser_Add(t *testing.T) {
	u := NewUser()

	req := &userPro.AddRequest{
		UniqueId: "258",
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

	req := &userPro.GetInfoByUserIdRequest{UserId:999999}

	res := &userPro.GetInfoByUserIdResponse{}

	err := u.GetInfoByUserId(context.TODO(), req, res)
	if err != nil {
		t.Error(err)
	}

	t.Log(res)
}

func TestUser_GetInfoByUniqueId(t *testing.T) {
	u := NewUser()

	req := &userPro.GetInfoByUniqueIdRequest{UniqueId:"258"}

	res := &userPro.GetInfoByUniqueIdResponse{}

	err := u.GetInfoByUniqueId(context.TODO(), req, res)
	if err != nil {
		t.Error(err)
	}

	t.Log(res.User)
}

func TestUser_GetListByUserId(t *testing.T) {
	userIds := []int64{1, 4, 999}

	u := NewUser()

	req := &userPro.GetListByUserIdRequest{UserIds:userIds}

	res := &userPro.GetListByUserIdResponse{}

	err := u.GetListByUserId(context.TODO(), req, res)
	if err != nil {
		t.Error(err)
	}

	t.Log(res)
}

func TestUser_CreateToken(t *testing.T) {
	var userId int64

	userId = 999

	u := NewUser()

	req := &userPro.CreateTokenRequest{UserId:userId}

	res := &userPro.CreateTokenResponse{}

	err := u.CreateToken(context.TODO(), req, res)
	if err != nil {
		t.Error(err)
	}

	t.Log(res)
}
