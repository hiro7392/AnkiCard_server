package crypto

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// 暗号(Hash)化
func PasswordEncrypt(password string) (string, error) {
	//hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	return string(hash), err
}

// 暗号(Hash)と入力された平パスワードの比較
func CompareHashAndPassword(hash, password string) error {
	fmt.Println("hash:", hash)
	fmt.Println("pass:", password)
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
