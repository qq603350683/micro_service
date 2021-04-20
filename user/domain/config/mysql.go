package config

import "github.com/micro/go-micro/v2/config"

/**
{
"host": "127.0.0.1",
"user": "root",
"password": "root",
"database": "test_micro",
"port": 3306
}
*/

type MysqlConfig struct {
	Host string `json:"host"`
	User string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
	Port int64 `json:"port"`
}

func NewMysqlConfig() *MysqlConfig {
	return new(MysqlConfig)
}

// 获取 Mysql 配置
func GetMysqlConfigFromConsul(config config.Config, path ...string) (*MysqlConfig, error) {
	c := NewMysqlConfig()

	err := config.Get(path...).Scan(c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
