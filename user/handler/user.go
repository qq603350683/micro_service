package handler

import (
	"context"
	"errors"
	"github.com/micro/go-micro/v2/util/log"
	"time"
	"user/domain/cache"
	"user/domain/model"
	"user/domain/repository"
	//"github.com/micro/go-micro/util/log"
	userProto "user/proto/user"
)

type User struct{
	DBService repository.IUserRepository
	CacheService cache.IUserCache
}

func NewUser() *User {
	u := new(User)
	u.DBService = repository.NewUserRepository()
	u.CacheService = cache.NewUserCache()

	return u
}

func (e *User) Add(ctx context.Context, req *userProto.AddRequest, res *userProto.AddResponse) error {
	var err error

	if req.UniqueId == "" {
		return errors.New("user.unique_id cannot be empty")
	}

	// 判断 req.UniqueId 是否已被注册
	ok, err := e.CacheService.CheckUniqueIDIsExists(req.UniqueId)
	if err != nil {
		return err
	}

	if ok {
		return errors.New("user.unique_id is exists")
	}

	user := model.NewUser()
	user.UniqueID = req.UniqueId

	err = e.DBService.Add(user)
	if err != nil {
		return err
	}

	ok, err = e.CacheService.AddUniqueID(req.UniqueId, user.UserId)
	if err != nil {
		return err
	}

	if !ok {
		return errors.New("user.unique_id cannot write in cache")
	}

	res.UserId = int64(user.UserId)

	return nil
}

func (e *User) GetInfoByUserID(ctx context.Context, req *userProto.GetInfoByUserIDRequest, res *userProto.GetInfoByUserIDResponse) error {
	var err error
	var user *model.User

	if req.UserId == 0 {
		return errors.New("user_id cannot be 0")
	}

	userId := int(req.UserId)

	// 1、先从缓存中获取
	user, err = e.CacheService.GetInfoByUserId(userId)
	if err != nil {
		return err
	}

	// 2、在缓存中获取失败再查找数据库
	if user == nil {
		// 先获取锁，避免发生并发
		key := cache.UserInfoLockKey(userId)
		ok, val := cache.Lock(key)
		if ok == true {
			user, err = e.DBService.GetInfoByUserId(userId, "")
			if err != nil {
				return nil
			}

			ok, err = e.CacheService.Add(userId, user)
			if err != nil {
				return err
			}

			if ok != true {
				log.Logf("user connot wirte in cache, user_id: %d", userId)
			}

			cache.UnLock(key, val)
		} else {
			// 2.1 睡眠一下再获取
			time.Sleep(time.Second / 5)
			user, err = e.CacheService.GetInfoByUserId(userId)
			if err != nil {
				return err
			}
		}
	}

	if user == nil {
		res.User = nil
	} else {
		res.User = &userProto.UserInfo{
			UserId:   int64(user.UserId),
			Nickname: user.Nickname,
			Avatar:   user.Avatar,
		}
	}

	return nil
}

// Call is a single request handler called via client.Call or the generated client code
//func (e *User) Call(ctx context.Context, req *user.Request, rsp *user.Response) error {
//	log.Log("Received User.Call request")
//	rsp.Msg = "Hello " + req.Name
//	return nil
//}

// Stream is a server side stream handler called via client.Stream or the generated client code
//func (e *User) Stream(ctx context.Context, req *user.StreamingRequest, stream user.User_StreamStream) error {
//	log.Logf("Received User.Stream request with count: %d", req.Count)
//
//	for i := 0; i < int(req.Count); i++ {
//		log.Logf("Responding: %d", i)
//		if err := stream.Send(&user.StreamingResponse{
//			Count: int64(i),
//		}); err != nil {
//			return err
//		}
//	}
//
//	return nil
//}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
//func (e *User) PingPong(ctx context.Context, stream user.User_PingPongStream) error {
//	for {
//		req, err := stream.Recv()
//		if err != nil {
//			return err
//		}
//		log.Logf("Got ping %v", req.Stroke)
//		if err := stream.Send(&user.Pong{Stroke: req.Stroke}); err != nil {
//			return err
//		}
//	}
//}
