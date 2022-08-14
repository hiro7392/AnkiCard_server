package service

import (
	"fmt"
	"log"

	"github.com/sakana7392/AnkiCard_server/application/crypto"
	//"github.com/sakana7392/AnkiCard_server/domain/model"
	"github.com/sakana7392/AnkiCard_server/infra/repository"
	"github.com/sakana7392/AnkiCard_server/testData"
)

// emaliとpasswordが存在するかチェック
func CheckEmailAndPassword(email, password string) (result bool, userId int, userName string) {

	//	結果とユーザ名を返す
	result = false
	userId = 0
	// emailが一致するユーザを取得
	User, err := repository.GetOneUserByEmail_DB(email)
	if err != nil {
		log.Println(err)
		return
	} else {
		fmt.Println("password:", password)
	}

	// passwordが一致するかチェック
	err = crypto.CompareHashAndPassword(User.Password, password)
	if err != nil {
		log.Println(err)
		fmt.Println(err)
		return
	}
	result = true
	userId = User.UserId
	userName = User.UserName
	return
}

// 新規作成時にemailとパスワードをチェック
func CreateNewUser() {

}

// テストデータをDBに挿入
func InsertTestUserData() {
	for _, user := range testData.TestUserData {
		repository.InserUserFromTestData_DB(user)
	}
	fmt.Println("testDateinserted successfully")
}
