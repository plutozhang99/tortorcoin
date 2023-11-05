package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword 用于散列密码
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash 用于验证密码是否与其散列版本匹配
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
