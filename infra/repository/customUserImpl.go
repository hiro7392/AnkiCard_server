package repository

import (
	"log"

	"github.com/sakana7392/AnkiCard_server/domain/model"
	"github.com/sakana7392/AnkiCard_server/infra"
)

// 1件取得
func GetAllCardsByUserId_DB(userId int) (cards []model.Card, err error) {

	rows, err := infra.Db.Query("SELECT card_id,tag_id,created_user_id,learning_level,question_text,answer_text FROM cards WHERE created_user_id=?", userId)
	for rows.Next() {
		var nowCard model.Card
		if err := rows.Scan(&nowCard.CardId, &nowCard.TagId, &nowCard.CreatedUserId, &nowCard.LearningLevel, &nowCard.QuestionText, &nowCard.AnswerText); err != nil {
			log.Fatal(err)
			log.Panicln(err)
		}
		cards = append(cards, nowCard)
	}

	return
}
