package repository

import (
	"fmt"
	"log"
	"time"

	"github.com/sakana7392/AnkiCard_server/domain/model"
	"github.com/sakana7392/AnkiCard_server/infra"
)

// 1件取得
func GetOneTag_DB(TagId int) (tag model.Tag, err error) {

	rows, err := infra.Db.Query("SELECT tag_id,creatted_user_id,tag_name FROM tags WHERE tag_id=?", TagId)
	for rows.Next() {
		if err := rows.Scan(&tag.TagId, &tag.CreatedUserId, &tag.TagName); err != nil {
			log.Fatal(err)
			log.Panicln(err)
		}
	}

	return
}

// 1件新規作成
func CreateNewTag_DB(tag *model.Tag) (err error) {
	var t = time.Now()
	const layout2 = "2006-01-02 15:04:05"

	_, err = infra.Db.Query("INSERT INTO tags(created_user_id,tag_name,created_at,updated_at) VALUES(?,?,?,?)",
		tag.CreatedUserId, tag.TagName, t.Format(layout2), t.Format(layout2))

	return err
}

// 1件削除
func DeleteOneTag_DB(tagId int) (err error) {
	_, err = infra.Db.Query("DELETE FROM tags WHERE tag_id=?", tagId)
	if err != nil {
		fmt.Println("deltefailed!")
		log.Panicln(err)
	}
	fmt.Println("delete success! tag_id=? deleted!", tagId)
	return
}

// 1件更新
func UpdateOneTag_DB(card *model.Card) (err error) {
	var t = time.Now()
	const layout2 = "2006-01-02 15:04:05"
	//更新される可能性があるのは、問題文、答え、タグIDのいずれか
	upd, err := infra.Db.Prepare("UPDATE cards SET tag_id=?,question_text=?,answer_text=?, updated_at=? WHERE card_id=?")
	if err != nil {
		fmt.Println("update failed! card_id=", card.CardId)
	}
	upd.Exec(card.TagId, card.QuestionText, card.AnswerText, t.Format(layout2), card.CardId)

	return
}
