package cache

import (
	"errors"
	"time"
	"user/domain/common"
	"user/domain/model"
)

type IUserTokenCache interface {
	Add(token *model.UserToken) (bool, error)
	Delete(token string) (bool, error)
	GetInfoByToken(token string) (*model.UserToken, error)
}

type UserTokenCache struct {

}

func NewUserTokenCache() IUserTokenCache {
	return new(UserTokenCache)
}

/**
 * @Description 新增token缓存
 * @param token UserToken详情
 */
func (u *UserTokenCache) Add(token *model.UserToken) (bool, error) {
	if token == model.NewUserToken() {
		return false, errors.New("token cannot be nil")
	}

	var err error

	key := UserTokenInfoKey(token.Token)

	_, err = RedisClient.HMSet(key, common.StructToMap(token)).Result()
	if err != nil {
		return false, err
	}

	// 设置过期时间
	diff := token.ExpiredAt.Sub(time.Now())
	if diff <= 0 {
		return false, errors.New("token.ExpiredAt cannot be expired")
	}

	_, err = RedisClient.Expire(key, diff).Result()
	if err != nil {
		return false, err
	}

	return true, nil
}

/**
 * @Description 删除token
 * @param token token字符串
 */
func (u *UserTokenCache) Delete(token string) (bool, error) {
	if token == "" {
		return false, errors.New("token is empty string")
	}

	key := UserTokenInfoKey(token)

	_, err := RedisClient.Del(key).Result()
	if err != nil {
		return false, err
	}

	return true, nil
}

/**
 * @Description 根据token获取详情
 * @param token token字符串
 */
func  (u *UserTokenCache) GetInfoByToken(token string) (*model.UserToken, error) {
	if token == "" {
		return nil, errors.New("token is empty string")
	}

	var err error

	key := UserTokenInfoKey(token)

	result, err := RedisClient.HGetAll(key).Result()
	if err != nil {
		return nil, err
	}

	if result == nil {
		return nil, nil
	}

	userToken := model.NewUserToken()

	err = common.MapToStruct(result, &userToken)
	if err != nil {
		return nil, err
	}

	if userToken == nil {
		return nil, nil
	}

	b := userToken.IsExpired()
	if b == true {
		_, err = u.Delete(token)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}

	return userToken, nil
}



