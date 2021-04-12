package repository

import "user/domain/model"

type IUserTokenRepository interface {
	Add(userToken *model.UserToken) error
	GetInfoByToken(token string) (*model.UserToken, error)
}

type UserTokenRepository struct {

}

func NewUserTokenRepository() IUserTokenRepository {
	return new(UserTokenRepository)
}

/**
 * @Description 新增一个token
 * @param userToken token详情
 */
func (u *UserTokenRepository) Add(userToken *model.UserToken) error {
	err := DB.Create(userToken).Error
	if err != nil {
		return err
	}

	return nil
}

/**
 * @Description  根据token返回token详情
 * @param token token
 */
func (u *UserTokenRepository) GetInfoByToken(token string) (*model.UserToken, error) {
	userToken := model.NewUserToken()

	err := DB.Where("token = ?", token).First(userToken).Error
	if err != nil {
		return nil, err
	}

	return userToken, nil
}
