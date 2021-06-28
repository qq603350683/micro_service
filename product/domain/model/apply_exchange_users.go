package model

import "time"

type ApplyExchangeUsers struct {
	ApplyExchangeUsersId int `json:"apply_exchange_users_id" gorm:"type:bigint(20) unsigned auto_increment;not null;primary_key"`
	UserId int `json:"user_id" gorm:"type:bigint(20) unsigned;not null;default:0;comment:'用户ID'"`
	ProductId int `json:"product_id" gorm:"type:bigint(20) unsigned;not null;default:0;comment:'产品ID'"`
	AppliedProdcut string `json:"applied_prodcut" gorm:"type:varchar(50);not null;default:'';comment:'申请产品'"`
	AppliedAt time.Time  `json:"applied_at" gorm:"column:applied_at"`
	Status bool `json:"status" gorm:"type:tinyint(1);not null;default:0;comment:'是否同意交换产品'"`

	Avatar string `json:"avatar" gorm:"-"`
	Nickname string `json:"nickname" gorm:"-"`
}