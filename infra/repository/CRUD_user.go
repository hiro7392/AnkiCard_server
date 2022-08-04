package repository

import (
	"log"

	"github.com/sakana7392/AnkiCard_server/domain"
	"github.com/sakana7392/AnkiCard_server/infra"
)

func GetOneUser_DB(userId int) (user domain.User, err error) {

	rows, err := infra.Db.Query("SELECT user_id,user_name,user_level FROM users WHERE user_id=?", userId)
	for rows.Next() {
		if err := rows.Scan(&user.UserId, &user.UserName, &user.UserLevel); err != nil {
			log.Fatal(err)
			log.Panicln(err)
		}
	}

	return
}
