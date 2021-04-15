package handler

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	"time"
	"webApi/common"
	lotteryProto "webApi/proto/lottery"

	//"github.com/golang/protobuf/proto"
	"github.com/micro/go-micro/v2"
	"net/http"
	webApi "webApi/proto/webApi"
	"webApi/response"
)

type WebApi struct{}

var Service micro.Service

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

	// 注册中心
	consulRegistry :=  consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.18.128:8500",
		}
	})

	send := &lotteryProto.AddRequest{
		Name:    name,
		BeginAt: beginAt,
		EndAt:   endAt,
	}

	s := micro.NewService()
	s.Init(micro.Registry(consulRegistry))
	cli := s.Client()

	// 这里很奇怪，好像没啥作用 time.Second * 30
	opts := cli.Options()
	opts.CallOptions.RequestTimeout = time.Second * 30
	opts.CallOptions.DialTimeout = time.Second * 30

	lotteryCli := lotteryProto.NewLotteryService("go.micro.service.lottery", cli)
	lotteryRes, err := lotteryCli.Add(context.TODO(), send)

	fmt.Println(lotteryRes)
	fmt.Println(err)

	res.Body = response.New("success", http.StatusOK, lotteryRes).ToString()

	return nil

	//var lotteryRes lotteryProto.AddResponse

	//opt := client.Option(func(options *client.Options) {
	//	options.Registry = consulRegistry
	//	options.ContentType = "application/grpc+proto"
	//})
	//
	//c := client.NewClient(opt)
	//
	//send := lotteryProto.AddRequest{
	//	Name:    "123123",
	//	BeginAt: "2021-03-08",
	//	EndAt:   "2021-03-18",
	//}
	//
	//b, err := proto.Marshal(&send)
	//if err != nil {
	//	panic(err)
	//}
	//
	//lotteryReq := c.NewRequest("go.micro.service.lottery", "Lottery.Add", b)
	//
	//err = client.Call(context.TODO(), lotteryReq, lotteryRes)
	//
	//fmt.Println(err)
}

func (e *WebApi) GetLotteryList(ctx context.Context, req *webApi.Request, res *webApi.Response) error {


	return nil

	//log.Info(req.Get[""])
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
