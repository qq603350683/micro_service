package cache

import (
	"testing"
	"time"
	"user/domain/model"
)

func TestUserTokenCache_Add(t *testing.T) {
	U := NewUserTokenCache()

	userToken := model.NewUserToken()
	userToken.Token = "test"
	userToken.ExpiredAt = time.Now().Add(5 * time.Second)
	userToken.userId = 1

	_, err := U.Add(userToken)
	if err != nil {
		t.Error(err)
	}
}

func TestUserTokenCache_Delete(t *testing.T) {
	U := NewUserTokenCache()

	b, err := U.Delete("test")
	if err != nil {
		t.Error(err)
	}
	t.Log(b)
}

func TestUserTokenCache_GetInfoByToken(t *testing.T) {
	U := NewUserTokenCache()

	userToken, err := U.GetInfoByToken("test")
	if err != nil {
		t.Error(err)
	}

	t.Log(userToken)
}
