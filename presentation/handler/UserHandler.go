package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"github.com/sakana7392/AnkiCard_server/domain/model"
	"github.com/sakana7392/AnkiCard_server/infra/repository"
)

func HandleUserRequest(w http.ResponseWriter, r *http.Request) {

	var err error
	switch r.Method {
	case "GET":
		err = GetOneUser(w, r)
	case "POST":
		err = CreateNewCard(w, r)
	case "PUT":
		err = UpdateOneCard(w, r)
	case "DELETE":
		err = DeleteOneCard(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	return
}

var privateCardFunc = http.HandlerFunc(HandleCardRequest)

// ユーザ情報を1件取得
func GetOneUser(w http.ResponseWriter, r *http.Request) (err error) {
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		log.Println(err)
		return
	}
	user, err := repository.GetOneUser_DB(id)
	if err != nil {
		log.Println(err)
		return
	}
	output, err := json.MarshalIndent(&user, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}

	w.Write(output)
	return
}

//	ユーザを新規作成
func CreateNewUser(w http.ResponseWriter, r *http.Request) (err error) {
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	u, _ := url.ParseQuery(r.URL.RawQuery)
	//fmt.Println(u)

	// query -> map[a:[AAA] b:[BBB] c:[CCC] d:[DDD]]

	var card model.Card

	card.AnswerText = u["answerText"][0]
	card.QuestionText = u["questionText"][0]
	fmt.Println(card.AnswerText)
	card.LearningLevel = 0
	card.TagId = 1

	err = repository.CreateNewCard_DB(&card)
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

//	既存のユーザを一件削除
func DeleteOneUser(w http.ResponseWriter, r *http.Request) (err error) {
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// query -> map[a:[AAA] b:[BBB] c:[CCC] d:[DDD]]
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		log.Println(err)
		return
	}
	err = repository.DeleteOneCard_DB(id)
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
func UpdateOneUser(w http.ResponseWriter, r *http.Request) (err error) {
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	u, _ := url.ParseQuery(r.URL.RawQuery)

	// query -> map[a:[AAA] b:[BBB] c:[CCC] d:[DDD]]

	var card model.Card

	card.AnswerText = u["answerText"][0]
	card.QuestionText = u["questionText"][0]
	card.CardId, err = strconv.Atoi(u["cardId"][0])
	if err != nil {
		log.Println(err)
		return
	}
	card.LearningLevel = 0
	card.TagId = 1

	err = repository.UpdateOneCard_DB(&card)
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
