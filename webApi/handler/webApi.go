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


// Call is a single request handler called via client.Call or the generated client code
//func (e *WebApi) Call(ctx context.Context, req *webApi.Request, rsp *webApi.Response) error {
//	log.Info("Received WebApi.Call request")
//	rsp.Msg = "Hello " + req.Name
//	return nil
//}

// Stream is a server side stream handler called via client.Stream or the generated client code
//func (e *WebApi) Stream(ctx context.Context, req *webApi.StreamingRequest, stream webApi.WebApi_StreamStream) error {
//	log.Infof("Received WebApi.Stream request with count: %d", req.Count)
//
//	for i := 0; i < int(req.Count); i++ {
//		log.Infof("Responding: %d", i)
//		if err := stream.Send(&webApi.StreamingResponse{
//			Count: int64(i),
//		}); err != nil {
//			return err
//		}
//	}
//
//	return nil
//}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
//func (e *WebApi) PingPong(ctx context.Context, stream webApi.WebApi_PingPongStream) error {
//	for {
//		req, err := stream.Recv()
//		if err != nil {
//			return err
//		}
//		log.Infof("Got ping %v", req.Stroke)
//		if err := stream.Send(&webApi.Pong{Stroke: req.Stroke}); err != nil {
//			return err
//		}
//	}
//}
