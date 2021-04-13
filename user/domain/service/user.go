package service

import (
	"errors"
	"github.com/micro/go-micro/v2/util/log"
	"time"
	"user/domain/cache"
	"user/domain/model"
	"user/domain/repository"
)

type IUserService interface {
	Add(*model.User) (int, error)
	GetInfoByUserId(userId int) (*model.User, error)
	GetInfoByUniqueId(uniqueId string) (*model.User, error)
	GetListByUserId(userIds []int) ([]model.User, error)
}

type UserService struct {
	DB repository.IUserRepository
	Cache cache.IUserCache
}

func NewUserService() IUserService {
	e := new(UserService)

	e.DB = repository.NewUserRepository()
	e.Cache = cache.NewUserCache()

	return e
}

/**
 * @Description 创建用户并生成缓存
 * @param user 用户详情
 * @return int 用户ID
 */
func (e *UserService) Add(user *model.User) (int, error) {
	// 判断 req.UniqueId 是否已被注册
	ok, err := e.Cache.CheckUniqueIDIsExists(user.UniqueId)
	if err != nil {
		return 0, err
	}

	if ok {
		return 0, errors.New("user.unique_id is exists")
	}

	err = e.DB.Add(user)
	if err != nil {
		return 0, err
	}

	ok, err = e.Cache.AddUniqueID(user.UniqueId, user.UserId)
	if err != nil {
		return 0, err
	}

	if !ok {
		return 0, errors.New("user.unique_id cannot write in cache")
	}

	return user.UserId, nil
}

/**
 * @Description 根据用户ID查找用户详情（先在缓存查找没有再到数据库查找）
 * @param userId 用户ID
 * @return *model.User 用户详情
 */
func (e *UserService) GetInfoByUserId(userId int) (*model.User, error) {
	var err error
	var user *model.User

	if userId <= 0 {
		return nil, errors.New("user_id cannot be 0")
	}

	// 1、先从缓存中获取
	user, err = e.Cache.GetInfoByUserId(userId)
	if err != nil {
		return nil, err
	}

	// 2、在缓存中获取失败再查找数据库
	if user == nil {
		// 先获取锁，避免发生并发
		key := cache.UserInfoLockKey(userId)
		ok, val := cache.Lock(key)
		if ok == true {
			user, err = e.DB.GetInfoByUserId(userId, "")
			if err != nil {
				return nil, nil
			}

			ok, err = e.Cache.Add(userId, user)
			if err != nil {
				return nil, err
			}

			if ok != true {
				log.Logf("user connot wirte in cache, user_id: %d", userId)
			}

			cache.UnLock(key, val)
		} else {
			// 2.1 睡眠一下再获取
			time.Sleep(time.Second / 5)
			user, err = e.Cache.GetInfoByUserId(userId)
			if err != nil {
				return nil, err
			}
		}
	} else {
		if user.DeletedAt != nil {
			return nil, nil
		}
	}

	return user, nil
}

/**
 * @Description 根据用户unique_id查找用户详情（先在缓存查找没有再到数据库查找）
 * @param userId 用户unique_id
 * @return *model.User 用户详情
 */
func (e *UserService) GetInfoByUniqueId(uniqueId string) (*model.User, error) {
	var ok bool
	var err error
	var userId int
	var user *model.User

	// 1、获取 unique_id 对应的 user_id
	userId, err = e.Cache.GetUserIdByUniqueID(uniqueId)
	if err != nil {
		return nil, err
	}

	if userId == -1 {
		// 这里返回-1是因为缓存key的field返回空值，为了避免缓存穿透做一层缓存
		ok, err = e.Cache.AddUniqueID(uniqueId, 0)
		if err != nil {
			return nil, err
		}

		if !ok {
			return nil, errors.New("user.unique_id cannot write in cache")
		}
	} else if (userId > 0) {
		user, err = e.GetInfoByUserId(userId)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}

/**
 * @Description 批量获取用户详情
 * @param userIds 用户ID切片
 * @return []model.User 用户详情切片
 */
func (e *UserService) GetListByUserId(userIds []int) ([]model.User, error) {
	var err error
	var users []model.User
	var cacheUsers []model.User

	if len(userIds) == 0 {
		return nil, nil
	}

	cacheUsers, err = e.Cache.GetListByUserId(userIds)
	if err != nil {
		return nil, err
	}

	if len(cacheUsers) == 0 {
		return nil, nil
	}

	for _, cacheUser := range(cacheUsers) {
		if (cacheUser.DeletedAt == nil) {
			users = append(users, cacheUser)
		}
	}

	if len(users) == 0 {
		return nil, nil
	}

	return users, nil
}

