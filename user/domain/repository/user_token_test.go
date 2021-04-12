package repository

import (
	"testing"
	"time"
	"user/domain/common"
	"user/domain/model"
)

func TestUserTokenRepository_Add(t *testing.T) {
	U := NewUserTokenRepository()

	userToken := model.NewUserToken()

	userToken.Token = common.CreateRandString(140, 150, "")
	userToken.UserId = 1
	userToken.ExpiredAt = time.Now().Add(24 * 3600 * 365 * time.Second)

	err := U.Add(userToken)
	if err != nil {
		t.Error(err)
	}
}

func TestUserTokenRepository_GetInfoByToken(t *testing.T) {
	U := NewUserTokenRepository()

	token := `ssyVWxuyBPnjUljBTnbdWxEojQlxDfURhXAJfTfuAOoTuIpntUhEJKdkHETZHOvdjZJfZy_rXMIuqEjACoJYurPYAZHX_MqRpSlmVlAOlcvZNPafic_hqlxPxZyHouiMHyxqopcuqvPSqtDjemEX`

	userToken, err := U.GetInfoByToken(token)
	if err != nil {
		t.Error(err)
	}

	t.Log(userToken)
}