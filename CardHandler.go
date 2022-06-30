package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"
	"strconv"
)

// カードを1件取得
func getOneCard(w http.ResponseWriter, r *http.Request) (err error) {
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		log.Println(err)
		return
	}
	card, err := getOneCard_DB(id)
	if err != nil {
		log.Println(err)
		return
	}
	output, err := json.MarshalIndent(&card, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}

	w.Write(output)
	return
}

//	カードを新規作成
func createNewCard(w http.ResponseWriter, r *http.Request) (err error) {
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	u, _ := url.ParseQuery(r.URL.RawQuery)
	//fmt.Println(u)

	// query -> map[a:[AAA] b:[BBB] c:[CCC] d:[DDD]]

	var card Card

	card.AnswerText = u["answerText"][0]
	card.QuestionText = u["questionText"][0]
	fmt.Println(card.AnswerText)
	card.LearningLevel = 0
	card.TagId = 1

	err = createNewCard_DB(&card)
	if err != nil {
		fmt.Println("failed to create new card")
		fmt.Println(err)
		w.WriteHeader(500)
		err.Error()
	} else {
		w.WriteHeader(200)
	}

	return err
}

//既存のカードを一件削除
func deleteOneCard(w http.ResponseWriter, r *http.Request) (err error) {
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// query -> map[a:[AAA] b:[BBB] c:[CCC] d:[DDD]]
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		log.Println(err)
		return
	}
	err = deleteOneCard_DB(id)
	if err != nil {
		fmt.Println("failed to create new card")
		fmt.Println(err)
		w.WriteHeader(500)
		err.Error()
	} else {
		w.WriteHeader(200)
	}

	return err
}

//	カード情報を更新
func updateOneCard(w http.ResponseWriter, r *http.Request) (err error) {
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	u, _ := url.ParseQuery(r.URL.RawQuery)

	// query -> map[a:[AAA] b:[BBB] c:[CCC] d:[DDD]]

	var card Card

	card.AnswerText = u["answerText"][0]
	card.QuestionText = u["questionText"][0]
	card.CardId, err = strconv.Atoi(u["cardId"][0])
	if err != nil {
		log.Println(err)
		return
	}
	card.LearningLevel = 0
	card.TagId = 1

	err = updateOneCard_DB(&card)
	if err != nil {
		fmt.Println("failed to update card")
		fmt.Println(err)
		w.WriteHeader(500)
		err.Error()
	} else {
		w.WriteHeader(200)
	}

	return err
}
