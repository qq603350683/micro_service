package common

import (
	"encoding/json"
	"math/rand"
	"time"
)

func StructToMap(s interface{}) map[string]interface{} {
	data, _ := json.Marshal(&s)

	m := make(map[string]interface{})

	json.Unmarshal(data, &m)

	return m
}

func MapToStruct(m interface{}, s interface{}) error {
	j, err := json.Marshal(m)
	if err != nil {
		return err
	}

	err = json.Unmarshal(j, &s)
	if err != nil {
		return err
	}

	return nil
}

// 创建随机字符串
func CreateRandString(min int, max int, chars string) string {
	if chars == "" {
		chars = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM_-"
	}

	length := 0
	if min == max && min == 0 {
		length = 100
	} else if min == max && min > 0 {
		length = min
	} else {
		rand.Seed(time.Now().Unix())
		length = rand.Intn(max - min) + min
	}

	str := ""
	char_len := len(chars)

	for length > 0 {
		rd := rand.New(rand.NewSource(time.Now().UnixNano()))
		i := rd.Intn(char_len - 1)

		str += chars[i:i+1]

		length -= 1

		time.Sleep(1000 * 0.001)
	}

	return str
}
