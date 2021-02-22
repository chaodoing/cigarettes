package utils

import (
	"crypto/md5"
	"fmt"

	uuid "github.com/satori/go.uuid"
)

// Uuid 生成唯一标识符
//	@return string key
func Uuid() (key string) {
	key = fmt.Sprintf("%v", uuid.NewV4())
	return
}

// Md5 加密字符串
//  @param value string 要加密的字符串
//  @return string  加密后的字符串
func Md5(value string) string {
	h := md5.Sum([]byte(value))
	md5String := fmt.Sprintf("%x", h)
	return md5String
}

// Password 加密密码
//  @param password string 要加密的密码
//  @return passwords string 加密后的密码
func Password(password string) (passwords string) {
	passwords = Md5(password + password + "chaodoing@live.com")
	return
}
