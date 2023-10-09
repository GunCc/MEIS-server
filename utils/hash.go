package utils

import "golang.org/x/crypto/bcrypt"

// 加密
func BcryptHash(val string) string {
	b, _ := bcrypt.GenerateFromPassword([]byte(val), bcrypt.DefaultCost)
	return string(b)
}

// 匹配
func BcryptCheck(password string, val string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(val))
	return err == nil
}
