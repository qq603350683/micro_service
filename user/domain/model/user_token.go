package model

import "time"

type UserToken struct {
	UserTokenId int `json:"user_token_id,string" gorm:"type:bigint(20) unsigned auto_increment;not null;primary_key"`
	UserId int `json:"user_id,string" gorm:"type:bigint(20) unsigned;not null;index:user_id_idx;comment:'用户ID 来自 user 表的 user_id'"`
	Token string `json:"token" gorm:"type:char(150);not null;default:'';unique_index;comment:'token'"`
	CreatedAt time.Time  `json:"-" gorm:"column:created_at;default:CURRENT_TIMESTAMP"`
	ExpiredAt time.Time  `json:"expired_at" gorm:"column:expired_at"`
}

func NewUserToken() *UserToken {
	return new(UserToken)
}

// token 进行 md5  主要用来缩短 redis 的键长度
func (u *UserToken) MD5Token() (string, error) {
	// TODO 这里 md5 加密测试暂时不做

	return u.Token, nil
}


func  (u *UserToken) IsExpired() bool {
	diff := u.ExpiredAt.Sub(time.Now())
	if diff <= 0 {
		return true
	} else {
		return false
	}
}