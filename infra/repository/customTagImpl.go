package repository

import (
	"log"

	"github.com/sakana7392/AnkiCard_server/domain/model"
	"github.com/sakana7392/AnkiCard_server/infra"
)

// 1件取得
func GetAllTagsByUserId_DB(userId int) (tags []model.Tag, err error) {

	rows, err := infra.Db.Query("select tag_id,tag_name,created_user_id from tags where created_user_id=?", userId)
	for rows.Next() {
		var nowTag model.Tag
		if err := rows.Scan(&nowTag.TagId, &nowTag.TagName, &nowTag.CreatedUserId); err != nil {
			log.Fatal(err)
			log.Panicln(err)
		}

		tags = append(tags, nowTag)
	}
	return
}
