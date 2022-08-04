package service

import (
	"fmt"
	"log"

	"github.com/sakana7392/AnkiCard_server/application/crypto"
	"github.com/sakana7392/AnkiCard_server/infra/repository"
	"github.com/sakana7392/AnkiCard_server/testData"
)

//	emaliとpasswordが存在するかチェック
func CheckEmailAndPassword(email, password string) bool {

	// emailが一致するユーザを取得
	User, err := repository.GetOneUserByEmail_DB(email)
	if err != nil {
		log.Println(err)
		return false
	}
	// フロントから送られてきたパスワードとDBに保存されているパスワードを比較
	cryptedPassword, err := crypto.PasswordEncrypt(password)
	if err != nil {
		log.Println(err)
		fmt.Println(err)
		return false
	}
	fmt.Println("cryptedPassword:", cryptedPassword)
	if crypto.CompareHashAndPassword(cryptedPassword,User.Password) != nil {
		log.Println(err)
		fmt.Println(err)
		return false
	}

	return true
}

//	新規作成時にemailとパスワードをチェック
func CreateNewUser() {

}

//	テストデータをDBに挿入
func InsertTestDate() {
	repository.InserUserFromTestData_DB(testData.TestUserData[0])
	repository.InserUserFromTestData_DB(testData.TestUserData[1])
	repository.InserUserFromTestData_DB(testData.TestUserData[2])
	fmt.Println("testDateinserted successfully")
}
