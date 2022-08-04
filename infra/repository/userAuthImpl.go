package repository

import (
	"log"

	"github.com/sakana7392/AnkiCard_server/domain/model"
	"github.com/sakana7392/AnkiCard_server/infra"
)

func GetOneUserByEmail_DB(userEmail,password string) (user model.User, err error) {

	rows, err := infra.Db.Query("SELECT user_id,user_name,user_level FROM users WHERE user_email=?", userEmail)
	for rows.Next() {
		if err := rows.Scan(&user.UserId, &user.UserName, &user.UserLevel); err != nil {
			log.Fatal(err)
			log.Panicln(err)
		}
	}


	return
}