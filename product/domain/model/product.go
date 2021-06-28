package model

type Product struct {
	ProductId int `json:"product_id" gorm:"type:bigint(20) unsigned auto_increment;not null;primary_key"`
	UserId int `json:"user_id" gorm:"type:bigint(20) unsigned;not null;default:0;comment:'用户ID'"`
	IsSpecifyCity bool `json:"is_specify_city" gorm:"type:tinyint(1);not null;default:0;comment:'是否指定城市'"`
	City string `json:"city" gorm:"type:varchar(10);not null;default:'';comment:'指定城市名称'"`
	MyProduct string `json:"my_product" gorm:"type:varchar(50);not null;default:'';comment:'我发布的产品'"`
	WantProduct string `json:"my_product" gorm:"type:varchar(50);not null;default:'';comment:'想要交换的产品'"`
	SuccessfullyExchangeUserId int `json:"user_id" gorm:"type:bigint(20) unsigned;not null;default:0;comment:'交换成功的用户ID'"`
	Status int8 `json:"is_specify_city" gorm:"type:tinyint(1);not null;default:0;comment:'状态 0 - 待交换 1 - 已交换成功'"`
	ProductImages []ProductImages `json:"product_images" gorm:"-"`
	ApplyExchangeUsers []ApplyExchangeUsers `json:"apply_exchange_users" gorm:"-"`
}

func NewProduct() *Product {
	return new(Product)
}