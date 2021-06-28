package model

type ProductImages struct {
	ProductImagesId int `json:"product_images_id" gorm:"type:bigint(20) unsigned auto_increment;not null;primary_key"`
	ProductId int `json:"product_id" gorm:"type:bigint(20) unsigned;not null;default:0;comment:'产品ID'"`
	Thumb string `json:"thumb" gorm:"type:varchar(50);not null;default:'';comment:'缩略图'"`
	Center string `json:"center" gorm:"type:varchar(50);not null;default:'';comment:'居中图'"`
	Original string `json:"original" gorm:"type:varchar(50);not null;default:'';comment:'原图'"`
}
 