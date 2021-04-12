package model

import (
	"testing"
	"time"
)

func TestUserToken_IsExpired(t *testing.T) {
	var b bool

	userToken := NewUserToken()

	userToken.ExpiredAt = time.Now().AddDate(0, 0, -2)
	b = userToken.IsExpired()
	if b != true {
		t.Errorf("当前时间为: %v, 过期时间为: %v, 结果应该 ture, 测试结果为 %t", time.Now(), userToken.ExpiredAt, b)
	}

	userToken.ExpiredAt = time.Now().AddDate(0, 0, 2)
	b = userToken.IsExpired()
	if b == true {
		t.Errorf("当前时间为: %v, 过期时间为: %v, 结果应该 false, 测试结果为 %t", time.Now(), userToken.ExpiredAt, b)
	}
}
