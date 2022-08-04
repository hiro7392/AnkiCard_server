package repository

import (
	"fmt"
	"log"
	"time"

	"github.com/sakana7392/AnkiCard_server/application/crypto"
	"github.com/sakana7392/AnkiCard_server/domain/model"
	"github.com/sakana7392/AnkiCard_server/infra"
)

func GetOneUserByEmail_DB(userEmail string) (user model.User, err error) {

	rows, err := infra.Db.Query("SELECT user_id,user_name,user_level FROM users WHERE user_email LIKE ?", userEmail)
	for rows.Next() {
		if err := rows.Scan(&user.UserId, &user.UserName, &user.UserLevel); err != nil {
			log.Fatal(err)
			log.Panicln(err)
		}
	}
	return
}

func InserUserFromTestData_DB(user model.User) (err error) {

	var t = time.Now()
	const layout2 = "2006-01-02 15:04:05"
	//passwordを暗号化
	user.Password, err = crypto.PasswordEncrypt(user.Password)
	_, err = infra.Db.Exec("INSERT INTO users (user_id,user_name,user_email,user_password,user_level,created_at,updated_at) VALUES (?,?,?,?,?,?,?)", user.UserId, user.UserName,user.Email, user.Password, user.UserLevel, t.Format(layout2), t.Format(layout2))
	if err != nil {
		log.Println(err)
		fmt.Println("insert failed!")
		return
	}
	return
}
