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
	GetInfoByToken(token string) (*model.UserToken, error)
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

/**
 * @Description: 根据 user_id 生成 token 令牌
 * @param userId 用户ID
 * @return *model.UserToken token令牌
 */
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

/**
 * @Description: 根据token获取令牌详情
 * @param token token
 * @return *model.UserToken 令牌详情
 */
func (e *UserTokenService) GetInfoByToken(token string) (*model.UserToken, error) {
	var err error
	var userToken *model.UserToken

	if token == "" {
		return nil, errors.New("token cannot be emtpy")
	}

	userToken, err = e.Cache.GetInfoByToken(token)
	if err != nil {
		return nil, err
	}

	if userToken == nil {
		return nil, nil
	}

	return userToken, nil
}