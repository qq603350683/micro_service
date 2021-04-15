package response

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/common/log"
)

type IResponse interface {
	ToString() string
}

type Response struct {
	Message string `json:"message"`
	Code int32 `json:"code"`
	Data interface{} `json:"data"`
	Test []string `json:"test"`
}

/**
 * @Description 序列化接口统一输出
 * @param message 返回给前端的信息
 * @param code 返回给前端的code
 * @param data 返回给前端的数据
 * @param opts 这里仅仅处理数组的第一个，主要用来调试用的
 * @return IResponse
 */
func New(message string, code int32, data interface{}, opts ...interface{}) IResponse {
	res := new(Response)

	res.Message = message
	res.Code = code
	res.Data = data

	if len(opts) > 0 {
		for _, opt := range(opts) {
			switch opt.(type) {
			case int:
				res.Test = append(res.Test, fmt.Sprintf("%v", opt))
			case string:
				res.Test = append(res.Test, fmt.Sprintf("%v", opt))
			}
		}
	}

	return res
}

func (r Response) ToString() string {
	body := make(map[string]interface{})

	body["message"] = r.Message
	body["code"] = r.Code
	body["data"] = r.Data

	if len(r.Test) > 0 {
		body["test"] = r.Test
	}

	b, err := json.Marshal(body)
	if err != nil {
		log.Info("JSON 解析失败")
		return ""
	}

	return string(b)
}
