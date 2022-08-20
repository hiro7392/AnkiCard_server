package repository

import (
	"log"

	"github.com/sakana7392/AnkiCard_server/domain/model"
	"github.com/sakana7392/AnkiCard_server/infra"
)

// 1件取得
func GetAllCardsByUserId_DB(userId int) (cards []model.Card, err error) {

	rows, err := infra.Db.Query("SELECT card_id,tag_id,tag_name,created_user_id,learning_level,question_text,answer_text FROM cards WHERE created_user_id=?", userId)
	for rows.Next() {
		var nowCard infra.NullableCard
		if err := rows.Scan(&nowCard.CardId, &nowCard.TagId, &nowCard.TagName, &nowCard.CreatedUserId, &nowCard.LearningLevel, &nowCard.QuestionText, &nowCard.AnswerText); err != nil {
			log.Fatal(err)
			log.Panicln(err)
		}
		card := model.Card{
			CardId:        nowCard.CardId,
			TagId:         int(nowCard.TagId.Int64),
			TagName:       nowCard.TagName.String,
			CreatedUserId: nowCard.CreatedUserId,
			LearningLevel: nowCard.LearningLevel,
			QuestionText:  nowCard.QuestionText,
			AnswerText:    nowCard.AnswerText,
		}
		cards = append(cards, card)
	}
	return
}
