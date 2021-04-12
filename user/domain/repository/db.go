package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
	"user/domain/model"
)

var DB *gorm.DB

func Database(connString string) {
	// connString 用户名:密码@(主机地址:端口)/数据库名称?charset=utf8&parseTime=True&loc=Local
	fmt.Println(connString)
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		panic(err)
	}

	// 禁止复表
	db.SingularTable(true)

	// 设置连接池
	// 空闲
	db.DB().SetMaxIdleConns(20)

	// 打开
	db.DB().SetMaxOpenConns(20)

	// 超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	// 开启SQL打印模式
	db.LogMode(true)

	DB = db

	migration()
}

func migration() {
	set := "ENGINE=InnoDB"

	user := model.NewUser()

	if DB.HasTable(user) {
		DB.AutoMigrate(user)
	} else {
		DB.Set("gorm:table_options", set).CreateTable(user)
	}

	userToken := model.NewUserToken()

	if DB.HasTable(userToken) {
		DB.AutoMigrate(userToken)
	} else {
		DB.Set("gorm:table_options", set).CreateTable(userToken)
	}
}
