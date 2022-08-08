package service

import (
	"fmt"
	"log"

	"github.com/sakana7392/AnkiCard_server/application/crypto"
	"github.com/sakana7392/AnkiCard_server/infra/repository"
	"github.com/sakana7392/AnkiCard_server/testData"
)

//	emaliとpasswordが存在するかチェック
func CheckEmailAndPassword(email, password string) (result bool, userName string) {

	//	結果とユーザ名を返す
	result=false
	userName=""
	// emailが一致するユーザを取得
	User, err := repository.GetOneUserByEmail_DB(email)
	if err != nil {
		log.Println(err)
		return
	}else{
		fmt.Println("password:", password)
	}
	
	// passwordが一致するかチェック
	err = crypto.CompareHashAndPassword(User.Password,password)
	if  err != nil {
		log.Println(err)
		fmt.Println(err)
		return
	}
	result=true
	userName=User.UserName
	return
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
