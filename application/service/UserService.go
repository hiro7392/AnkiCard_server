package service

import (
	"fmt"
	"log"

	"github.com/sakana7392/AnkiCard_server/application/crypto"
	"github.com/sakana7392/AnkiCard_server/infra/repository"
)

//	emaliとpasswordが存在するかチェック
func CheckEmailAndPassword(email, password string) bool {

	// emailが一致するユーザを取得
	User, err := repository.GetOneUserByEmail_DB(email, password)
	if err != nil {
		log.Println(err)
		return false
	}
	// フロントから送られてきたパスワードとDBに保存されているパスワードを比較
	cryptedPassword, err := crypto.PasswordEncrypt(password)
	if crypto.CompareHashAndPassword(cryptedPassword, User.Password) != nil {
		log.Println(err)
		fmt.Println(err)
		return false
	}

	return true
}

//	新規作成時にemailとパスワードをチェック
func CreateNewUser() {

}
