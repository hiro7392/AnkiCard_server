package main

import (
	"fmt"
	"log"
	"time"
)

// 1件取得
func getOneCard_DB(cardId int) (card Card,err error) {

	rows, err := db.Query("SELECT card_id,question_text,answer_text FROM cards WHERE card_id=?",cardId)
	for rows.Next(){
		if err:=rows.Scan(&card.CardId,&card.QuestionText,&card.AnswerText); err!=nil{
			log.Fatal(err)
			log.Panicln(err)
		}
	}

	return
}

//	1件新規作成
func createNewCard_DB(card *Card) (err error) {
	var t = time.Now()
	const layout2 = "2006-01-02 15:04:05"

	_, err = db.Query("INSERT INTO cards(card_id,tag_id,learning_level,question_text,answer_text,created_at,updated_at) VALUES(?,?,?,?,?,?,?)",
		card.CardId, card.TagId, card.LearningLevel, card.QuestionText, card.AnswerText, t.Format(layout2), t.Format(layout2))

	return err
}

// 1件削除
func deleteOneCard_DB(cardId int) (err error) {
	_, err = db.Query("DELETE FROM cards WHERE card_id=?", cardId)
	if err != nil {
		fmt.Println("deltefailed!")
		log.Panicln(err)
	}
	fmt.Println("delete success! card_id=? deleted!", cardId)
	return
}

//1件更新
func updateOneCard_DB(card *Card) (err error) {

	var t = time.Now()
	const layout2 = "2006-01-02 15:04:05"
	//更新される可能性があるのは、問題文、答え、タグIDのいずれか
	upd, err := db.Prepare("UPDATE cards SET tag_id=?,question_text=?,answer_text=?, updated_at=? WHERE card_id=?")
	if err != nil {
		fmt.Println("update failed! card_id=", card.CardId)
	}
	upd.Exec(card.TagId, card.QuestionText, card.AnswerText, t.Format(layout2), card.CardId)

	return
}
