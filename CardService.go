package main

import (
	"fmt"
	"net/http"
	"net/url"
)

//	カードを新規作成
func createNewCard(w http.ResponseWriter, r *http.Request) (err error) {
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	u, _ := url.ParseQuery(r.URL.RawQuery)
	fmt.Println(u)

	// query -> map[a:[AAA] b:[BBB] c:[CCC] d:[DDD]]

	var card Card
	card.AnswerText = u["answerText"][0]
	card.QuestionText = u["questionText"][0]
	card.LearningLevel = 0
	card.TagId = 1

	card.createNewCard_DB()
	w.WriteHeader(200)
	return
}
