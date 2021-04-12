package service

import (
	"errors"
	"time"
	"user/domain/cache"
	"user/domain/common"
	"user/domain/model"
	"user/domain/repository"
)

type IUserTokenService interface {
	CreateToken(userId int) (*model.UserToken, error)
	//GetUserInfoByToken(token string) (*model.UserToken, error)
}

type UserTokenService struct {
	DB repository.IUserTokenRepository
	Cache cache.IUserTokenCache
}

func NewUserTokenService() IUserTokenService {
	e := new(UserTokenService)

	e.DB = repository.NewUserTokenRepository()
	e.Cache = cache.NewUserTokenCache()

	return e
}

func (e *UserTokenService) CreateToken(userId int) (*model.UserToken, error) {
	var err error
	var ok bool

	if userId == 0 {
		return nil, errors.New("user_id cannot be 0")
	}

	userToken := model.NewUserToken()
	userToken.UserId = userId
	userToken.Token = common.CreateRandString(140, 150, "")
	userToken.ExpiredAt = time.Now().Add(time.Second * (3600 * 24 * 365))

	err = e.DB.Add(userToken)
	if err != nil {
		return nil, err
	}

	ok, err = e.Cache.Add(userToken)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, errors.New("系统繁忙，请稍后再试")
	}

	return userToken, nil
}