package repository

import (
	"errors"
	"github.com/jinzhu/gorm"
	"user/domain/model"
)

type IUserRepository interface {
	Add(user *model.User) error
	Update(user *model.User) error
	Delete(userId int) error
	GetInfoByUniqueID(uniqueID string, fields string) (*model.User, error)
	GetInfoByUserId(userId int, fields string) (*model.User, error)
	GetListByUserId(userId []int, fields string) ([]model.User, error)
}

type UserRepository struct {

}

func NewUserRepository() IUserRepository {
	return new(UserRepository)
}

/**
 * @Description 新增一个用户
 * @param user 用户详情
 */
func (i *UserRepository) Add(user *model.User) error {
	err := DB.Create(user).Error

	return err
}

/**
 * @Description 修改用户信息
 * @param lottery 用户详情
 */
func (i *UserRepository) Update(user *model.User) error {
	res := DB.Model(user).Update(user).RowsAffected
	if res == 0 {
		return errors.New("更新失败")
	}

	return nil
}

/**
 * @Description  删除用户
 * @param userId 用户ID
 */
func (i *UserRepository) Delete(userId int) error {
	res := DB.Where("user_id = ?", userId).Delete(model.NewUser()).RowsAffected
	if res == 0 {
		return errors.New("删除失败")
	}

	return nil
}

/**
 * @Description  根据openid获取用户详情
 * @param uniqueID 唯一ID
 * @param fields 筛选的字段
 * @return *model.User 用户详情
 */
func (i *UserRepository) GetInfoByUniqueID(uniqueID string, fields string) (*model.User, error) {
	user := model.NewUser()

	if fields == "" {
		fields = "*"
	}

	err := DB.Select(fields).Where("unique_id = ?", uniqueID).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}

/**
 * @Description  根据user_id获取用户详情
 * @param openid openid
 * @param fields 筛选的字段
 * @return *model.User 用户详情
 */
func (i *UserRepository) GetInfoByUserId(userId int, fields string) (*model.User, error) {
	user := model.NewUser()

	if fields == "" {
		fields = "*"
	}

	err := DB.Select(fields).Where("user_id = ?", userId).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return user, nil
}


/**
 * @Description  批量获取用户信息
 * @param userId 用户ID数组
 * @param fields 筛选的字段
 * @return []*model.User 用户详情
 */
func (i *UserRepository) GetListByUserId(userId []int, fields string) ([]model.User, error) {
	userIdLen := len(userId)
	if userIdLen == 0 {
		return nil, nil
	}

	if fields == "" {
		fields = "*"
	}

	var users []model.User

	err := DB.Select(fields).Where("user_id IN (?)", userId).Find(&users).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
	}

	return users, nil
}

