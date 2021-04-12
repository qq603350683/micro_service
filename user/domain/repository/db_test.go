package repository

import "testing"

func TestMain(m *testing.M) {
	Database("root:root@(127.0.0.1:3306)/test_micro?charset=utf8&parseTime=True&loc=Local")

	m.Run()
}
