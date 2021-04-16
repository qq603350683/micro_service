package handler

import (
	"context"
	"strconv"
	"webApi/common"
	"webApi/domain/service"
	lotteryProto "webApi/proto/lottery"
	userProto "webApi/proto/user"

	"net/http"
	webApi "webApi/proto/webApi"
	"webApi/response"
)

type WebApi struct{}

//var Service micro.Service

type LotteryResponse struct {
	LotteryID int64
}

// http://127.0.0.1:8080/webApi/addLottery
func (e *WebApi) AddLottery(ctx context.Context, req *webApi.Request, res *webApi.Response) error {
	res.StatusCode = http.StatusOK

	var ok bool
	var name string
	var beginAt string
	var endAt string

	if _, ok = req.Post["name"]; !ok {
		res.Body = response.New("请输入奖品标题", http.StatusInternalServerError, nil).ToString()
		return nil
	} else {
		name = req.Post["name"].Values[0]
		nameLen := len(name)
		if nameLen == 0 {
			res.Body = response.New("请输入奖品标题", http.StatusInternalServerError, nil).ToString()
			return nil
		}

		if nameLen > 20 {
			res.Body = response.New("奖品标题最多输入20个字", http.StatusInternalServerError, nil).ToString()
			return nil
		}

		// 过滤XSS攻击
		name = common.XssFilter(name)
	}

	if _, ok = req.Post["begin_at"]; !ok {
		res.Body = response.New("请选择奖品开始抽奖时间", http.StatusInternalServerError, nil).ToString()
		return nil
	} else {
		beginAt = req.Post["begin_at"].Values[0]
		if common.IsDate(beginAt) == false {
			res.Body = response.New("请选择正确的奖品开始抽奖时间", http.StatusInternalServerError, nil).ToString()
			return nil
		}
	}

	if _, ok = req.Post["end_at"]; !ok {
		res.Body = response.New("请选择奖品结束抽奖时间", http.StatusInternalServerError, nil).ToString()
		return nil
	} else {
		endAt = req.Post["begin_at"].Values[0]
		if common.IsDate(endAt) == false {
			res.Body = response.New("请选择正确的奖品结束抽奖时间", http.StatusInternalServerError, nil).ToString()
			return nil
		}
	}

	params := &lotteryProto.AddRequest{
		Name:    name,
		BeginAt: beginAt,
		EndAt:   endAt,
	}

	cli := lotteryProto.NewLotteryService(service.ServiceLottery, service.GetClient())

	lotteryRes, err := cli.Add(context.TODO(), params)
	if err != nil {
		return err
	}

	res.Body = response.New("success", http.StatusOK, lotteryRes).ToString()

	return nil
}

func (e *WebApi) GetLotteryList(ctx context.Context, req *webApi.Request, res *webApi.Response) error {
	return nil
}

func (e *WebApi) GetInfoByUserId(ctx context.Context, req *webApi.Request, res *webApi.Response) error {
	res.StatusCode = http.StatusOK

	var ok bool
	var err error
	var userId int

	if _, ok = req.Get["user_id"]; !ok {
		res.Body = response.New("缺少user_id参数", http.StatusInternalServerError, nil).ToString()
		return nil
	} else {
		str := req.Get["user_id"].Values[0]
		userId, err = strconv.Atoi(str)
		if err != nil {
			return err
		}
	}

	if userId == 0 {
		return nil
	}

	params := &userProto.GetInfoByUserIdRequest{UserId:int64(userId)}

	cli := userProto.NewUserService(service.ServiceUser, service.GetClient())

	userRes, err := cli.GetInfoByUserId(context.TODO(), params)
	if err != nil {
		return err
	}

	if userRes == nil {
		res.Body = response.New("success", http.StatusOK, nil).ToString()
	} else {
		res.Body = response.New("success", http.StatusOK, userRes.User).ToString()
	}

	return nil

}

func (e *WebApi) GetListByUserId(ctx context.Context, req *webApi.Request, res *webApi.Response) error {
	var ok bool
	var userIds []int64

	if _, ok = req.Post["user_ids"]; !ok {
		res.Body = response.New("缺少user_ids参数", http.StatusInternalServerError, nil, "缺少user_ids参数").ToString()
		return nil
	}

	if len(req.Post["user_ids"].Values) == 0 {
		res.Body = response.New("user_ids 不能是空数组", http.StatusInternalServerError, nil, "user_ids 不能是空数组").ToString()
		return nil
	}

	for _, value := range(req.Post["user_ids"].Values) {
		userId, err := strconv.Atoi(value)
		if err != nil {
			return err
		}

		userIds = append(userIds, int64(userId))
	}

	params := &userProto.GetListByUserIdRequest{UserIds:userIds}

	cli := userProto.NewUserService(service.ServiceUser, service.GetClient())

	list, err := cli.GetListByUserId(context.TODO(), params)
	if err != nil {
		return err
	}

	if list == nil {
		res.Body = response.New("success", http.StatusOK, nil).ToString()
	} else {
		res.Body = response.New("success", http.StatusOK, list.Users).ToString()
	}

	return nil
}

func (e *WebApi) Login(ctx context.Context, req *webApi.Request, res *webApi.Response) error {
	var ok bool
	var uniqueId string
	
	if _, ok = req.Post["unique_id"]; !ok {
		res.Body = response.New("请输入账户", http.StatusInternalServerError, nil, "缺少unique_id参数").ToString()
		return nil
	}

	uniqueId = req.Post["unique_id"].Values[0]
	if uniqueId == "" {
		res.Body = response.New("请输入账户", http.StatusInternalServerError, nil, "账户不能为空").ToString()
		return nil
	}

	params := &userProto.GetInfoByUniqueIdRequest{UniqueId:uniqueId}

	cli := userProto.NewUserService(service.ServiceUser, service.GetClient())

	resp, err := cli.GetInfoByUniqueId(context.TODO(), params)
	if err != nil {
		return err
	}

	if resp == nil {
		res.Body = response.New("账户错误", http.StatusInternalServerError, nil).ToString()
		return nil
	}

	if resp.User == nil {
		res.Body = response.New("账户错误", http.StatusInternalServerError, nil).ToString()
		return nil
	}

	param1 := &userProto.CreateTokenRequest{UserId:int64(resp.User.UserId)}

	resp1, err := cli.CreateToken(context.TODO(), param1)
	if err != nil {
		return err
	}

	res.Body = response.New("账户错误", http.StatusInternalServerError, resp1).ToString()

	return nil
}

func (e *WebApi) Register(ctx context.Context, req *webApi.Request, res *webApi.Response) error {
	var ok bool
	var uniqueId string

	if _, ok = req.Post["unique_id"]; !ok {
		res.Body = response.New("请输入账户", http.StatusInternalServerError, nil, "缺少unique_id参数").ToString()
		return nil
	}

	uniqueId = req.Post["unique_id"].Values[0]
	if uniqueId == "" {
		res.Body = response.New("请输入账户", http.StatusInternalServerError, nil, "账户不能为空").ToString()
		return nil
	}

	param := &userProto.GetInfoByUniqueIdRequest{UniqueId:uniqueId}

	cli := userProto.NewUserService(service.ServiceUser, service.GetClient())

	resp, err := cli.GetInfoByUniqueId(context.TODO(), param)
	if err != nil {
		return err
	}

	if resp != nil && resp.User != nil {
		res.Body = response.New("当前账户已被注册", http.StatusInternalServerError, nil).ToString()
		return nil
	}

	return nil
	//param1 := &userProto.A
}
