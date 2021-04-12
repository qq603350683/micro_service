package cache

import "fmt"

func UserInfoLockKey(userId int) string {
	return fmt.Sprintf("user:info:lock:%d", userId)
}

func UserInfoKey(userId int) string {
	return fmt.Sprintf("user:info:%d", userId)
}

func UserUniqueIDListKey() string {
	return fmt.Sprintf("user:unique_id:list")
}

func UserTokenInfoKey(token string) string {
	return fmt.Sprintf("user_token:info:%s", token)
}
