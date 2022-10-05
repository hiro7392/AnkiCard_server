package repository

import (
	"fmt"
	"log"
	"time"

	"github.com/sakana7392/AnkiCard_server/domain/model"
	"github.com/sakana7392/AnkiCard_server/infra"
)

// 1件取得
func GetOneCard_DB(cardId int) (card model.Card, err error) {

	rows, err := infra.Db.Query("SELECT card_id,tag_id,created_user_id,learning_level,question_text,answer_text FROM cards WHERE card_id=?", cardId)
	for rows.Next() {
		if err := rows.Scan(&card.CardId, &card.TagId, &card.CreatedUserId, &card.LearningLevel, &card.QuestionText, &card.AnswerText); err != nil {
			log.Fatal(err)
			log.Panicln(err)
		}
	}

	return
}

// 1件新規作成
func CreateNewCard_DB(card *model.Card) (err error) {
	var t = time.Now()
	const layout2 = "2006-01-02 15:04:05"

	_, err = infra.Db.Query("INSERT INTO cards(tag_id,tag_name,created_user_id,learning_level,question_text,answer_text,created_at,updated_at) VALUES(?,?,?,?,?,?,?,?)",
		card.TagId, card.TagName, card.CreatedUserId, card.LearningLevel, card.QuestionText, card.AnswerText, t.Format(layout2), t.Format(layout2))

	return err
}

// 1件削除
func DeleteOneCard_DB(cardId int) (err error) {
	_, err = infra.Db.Query("DELETE FROM cards WHERE card_id=?", cardId)
	if err != nil {
		fmt.Println("deltefailed!")
		log.Panicln(err)
	}
	fmt.Println("delete success! card_id=? deleted!", cardId)
	return
}

// 1件更新
func UpdateOneCard_DB(card *model.Card) (err error) {
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

func UpdateOneCardLevel_DB(addLevel, cardId int) (existCard model.Card, err error) {

	// まずは既存のカードを取得
	existCard, err = GetOneCard_DB(cardId)
	if err != nil {
		fmt.Println("getting card failed! card_id=", cardId)
	}
	_, err = infra.Db.Query("UPDATE cards SET learning_level = learning_level + ?  WHERE card_id = ?", addLevel, cardId)
	if err != nil {
		fmt.Println("update failed! card_id=", cardId)
		return existCard, err
	}
	return
}
