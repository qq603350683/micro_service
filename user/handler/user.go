package handler

import (
	"context"
	"errors"
	"user/domain/model"
	"user/domain/service"
	userProto "user/proto/user"
)

type User struct{
	UserService service.IUserService
	UserTokenService service.IUserTokenService
}

func NewUser() *User {
	e := new(User)

	e.UserService = service.NewUserService()
	e.UserTokenService = service.NewUserTokenService()

	return e
}

// 添加用户详情
func (e *User) Add(ctx context.Context, req *userProto.AddRequest, res *userProto.AddResponse) error {
	var err error

	if req.UniqueId == "" {
		return errors.New("user.unique_id cannot be empty")
	}

	user := model.NewUser()
	user.UniqueId = req.UniqueId
	user.Nickname = req.Nickname
	user.Avatar = req.Avatar

	user.UserId, err = e.UserService.Add(user)
	if err != nil {
		return err
	}

	res.UserId = int64(user.UserId)

	return nil
}

// 根据用户ID获取用户详情
func (e *User) GetInfoByUserId(ctx context.Context, req *userProto.GetInfoByUserIdRequest, res *userProto.GetInfoByUserIdResponse) error {
	var err error
	var user *model.User

	if req.UserId == 0 {
		return errors.New("user_id cannot be 0")
	}

	userId := int(req.UserId)

	user, err = e.UserService.GetInfoByUserId(userId)
	if err != nil {
		return nil
	}

	if user != nil {
		res.User = &userProto.UserInfo{
			UserId:   int64(user.UserId),
			Nickname: user.Nickname,
			Avatar:   user.Avatar,
		}
	}

	return nil
}

// 根据用户unique_id获取用户详情
func (e *User) GetInfoByUniqueId(ctx context.Context, req *userProto.GetInfoByUniqueIdRequest, res *userProto.GetInfoByUniqueIdResponse) error {
	if req.UniqueId == "" {
		return errors.New("unique_id cannot be empty")
	}

	user, err := e.UserService.GetInfoByUniqueId(req.UniqueId)
	if err != nil {
		return err
	}

	if user != nil {
		res.User = &userProto.UserInfo{
			UserId:   int64(user.UserId),
			Nickname: user.Nickname,
			Avatar:   user.Avatar,
		}
	}

	return nil
}

// 根据 user_id 批量获取 用户信息
func (e *User) GetListByUserId(ctx context.Context, req *userProto.GetListByUserIdRequest, res *userProto.GetListByUserIdResponse) error {
	var err error
	var users []model.User

	if len(req.UserIds) == 0 {
		return nil
	}

	var userIds []int

	for _, userId := range(req.UserIds) {
		userIds = append(userIds, int(userId))
	}

	users, err = e.UserService.GetListByUserId(userIds)
	if err != nil {
		return err
	}

	if users != nil {
		for _, user := range(users) {
			res.Users = append(res.Users, &userProto.BaseUser{
				UserId:   int64(user.UserId),
				Nickname: user.Nickname,
				Avatar:   user.Avatar,
			})
		}
	}

	return nil
}

// 根据用户ID生成token
func (e *User) CreateToken(ctx context.Context, req *userProto.CreateTokenRequest, res *userProto.CreateTokenResponse) error {
	var err error
	var userId int
	var token *model.UserToken
	var user *model.User

	if req.UserId == 0 {
		return errors.New("user_id cannot be 0")
	}

	userId = int(req.UserId)

	// 判断用户是否存在
	user, err = e.UserService.GetInfoByUserId(userId)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("用户不存在")
	}

	token, err = e.UserTokenService.CreateToken(userId)
	if err != nil {
		return err
	}

	res.Token = token.Token
	res.ExpiredAt = token.ExpiredAt.Format("2006-01-02 15:04:05")

	return nil
}