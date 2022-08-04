package service

import (
	"fmt"
	"log"

	"github.com/sakana7392/AnkiCard_server/crypto"
	"github.com/sakana7392/AnkiCard_server/infra/repository"
)

//	emaliとpasswordが存在するかチェック
func checkEmailAndPassword(email, password string) bool {

	// emailが一致するユーザを取得
	User, err := repository.GetOneUserByEmail_DB(email, password)
	if err != nil {
		log.Println(err)
		return false
	}
	// フロントから送られてきたパスワードとDBに保存されているパスワードを比較
	cryptedPassword := crypto.cryptPassword(password)
	if (crypto.ComparePassword(cryptedPassword, User.Password)!=nil) {
		log.Println(err)
		fmt.Println(err)
	}

	return true
}

//	新規作成時にemailとパスワードをチェック
func createNewUser() {

}
