package cache

import (
	"github.com/go-redis/redis"
	"log"
	"strconv"
	"time"
	"user/domain/common"
	"user/domain/model"
)

type IUserCache interface {
	Add(userId int, user *model.User) (bool, error)
	Update(userId int, user *model.User) (bool, error)
	Delete(userId int) (bool, error)
	GetInfoByUserId(userId int) (*model.User, error)
	// GetInfoByUniqueID(uniqueID string) (*model.User, error)
	GetUserIdByUniqueID(uniqueID string) (int, error)
	AddUniqueID(uniqueID string, userId int) (bool, error)
	CheckUniqueIDIsExists(uniqueID string) (bool, error)
	GetListByUserId(userIds []int) ([]model.User, error)
}

type UserCache struct {

}

func NewUserCache() IUserCache {
	return new(UserCache)
}

// 添加用户详情缓存
func (u *UserCache) Add(userId int, user *model.User) (bool, error) {
	var err error

	key := UserInfoKey(userId)

	if user == nil || user.DeletedAt != nil {
		// 这里是空值缓存，主要用来防止缓存穿透

		_, err = RedisClient.HMSet(key, model.NewUserByuserId(userId).ToEmptyCache()).Result()
		if err != nil {
			return false, err
		}

		// 设置一个过期时间，让它一段时间后消失，避免一直占用缓存
		_, err = RedisClient.Expire(key, time.Second * 300).Result()
		if err != nil {
			return false, err
		}
	} else {
		// 缓存
		_, err := RedisClient.HMSet(key, common.StructToMap(user)).Result()
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

// 修改用户详情缓存
func (u *UserCache) Update(userId int, user *model.User) (bool, error) {
	return u.Add(userId, user)
}

// 删除缓存
func (u *UserCache) Delete(userId int) (bool, error) {
	// 设置为nil值，让缓存自动回收删除，防止删除时候并发查库再写入空值
	var err error

	key := UserInfoKey(userId)

	_, err = RedisClient.HMSet(key, model.NewUserByuserId(userId).ToEmptyCache()).Result()
	if err != nil {
		return false, err
	}

	_, err = RedisClient.Expire(key, time.Second * 300).Result()
	if err != nil {
		return false, err
	}

	return true, nil
}

// 根据用户ID获取详情
func (u *UserCache) GetInfoByUserId(userId int) (*model.User, error) {
	log.Printf("Cache: GetInfoByUserId(%d)", userId)

	var err error
	var user *model.User
	var result = make(map[string]string)

	key := UserInfoKey(userId)

	result, err = RedisClient.HGetAll(key).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		} else {
			return nil, err
		}
	}

	err = common.MapToStruct(result, &user)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, nil
	}

	if user.DeletedAt != nil {
		return nil, nil
	}

	return user, nil
}

// 根据 unique_id 获取 user_id, 当 user_id 为 -1 代表当前 unique_id 不在缓存里
func (u *UserCache) GetUserIdByUniqueID(uniqueID string) (int, error) {
	key := UserUniqueIDListKey()

	userIdStr, err := RedisClient.HGet(key, uniqueID).Result()
	if err != nil {
		return 0, err
	}

	if userIdStr == "" {
		return -1, nil
	}

	return strconv.Atoi(userIdStr)
}

// 根据 unique_id 获取详情
//func (u *UserCache) GetInfoByUniqueID(uniqueID string) (*model.User, error) {
//	key := UserUniqueIDListKey()
//
//	str, err := RedisClient.HGet(key, uniqueID).Result()
//	if err != nil && err != redis.Nil {
//		return nil, err
//	}
//
//	if str == "" {
//		return nil, nil
//	}
//
//	userId, err := strconv.Atoi(str)
//	if err != nil {
//		return nil, nil
//	}
//
//	user, err := u.GetInfoByUserId(userId)
//
//	return user, err
//}

// 把openid添加到数组
func (u *UserCache) AddUniqueID(uniqueID string, userId int) (bool, error) {
	key := UserUniqueIDListKey()

	_, err := RedisClient.HSet(key, uniqueID, userId).Result()
	if err != nil {
		return false, nil
	}

	return true, nil
}

// 检测openid是否存在
func (u *UserCache) CheckUniqueIDIsExists(uniqueID string) (bool, error) {
	key := UserUniqueIDListKey()

	return RedisClient.HExists(key, uniqueID).Result()
}

// 批量获取用户详情
func (u *UserCache) GetListByUserId(userIds []int) ([]model.User, error) {
	var err error
	var users []model.User

	userIdLen := len(userIds)
	if userIdLen == 0 {
		return nil, nil
	}
	pipe := RedisClient.Pipeline()

	for _, userId := range(userIds) {
		key := UserInfoKey(userId)

		pipe.HGetAll(key)
	}

	cmders, err := pipe.Exec()
	if err != nil {
		if err == redis.Nil {
			return nil, nil
		}
		return nil, err
	}

	for _, cmder := range(cmders) {
		result, err := cmder.(*redis.StringStringMapCmd).Result()
		if err != nil {
			if err == redis.Nil {
				continue
			}
			return nil, err
		}

		// 如果不存在会返回 map[]， 判断一下长度，如果是 0 则是没有任何内容
		if len(result) == 0 {
			continue
		}
		var user model.User

		err = common.MapToStruct(result, &user)
		if err != nil {
			return nil, err
		}

		if user.DeletedAt == nil {
			users = append(users, user)
		}
	}

	if len(users) == 0 {
		return nil, nil
	}

	return users, nil
}