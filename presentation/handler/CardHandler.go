package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sakana7392/AnkiCard_server/domain/model"
	"github.com/sakana7392/AnkiCard_server/infra/repository"
)

func HandleCardRequest(w http.ResponseWriter, r *http.Request) {

	fmt.Println("HandlerCardRequest")
	var err error
	switch r.Method {
	case "GET":
		err = GetOneCard(w, r)
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

// カードを1件取得
func GetOneCard(w http.ResponseWriter, r *http.Request) (err error) {
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	vars := mux.Vars(r)
	id := vars["id"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		return
	}
	card, err := repository.GetOneCard_DB(idInt)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("result=", card)
	output, err := json.MarshalIndent(&card, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}

	w.Write(output)
	return
}

// カードを新規作成
func CreateNewCard(w http.ResponseWriter, r *http.Request) (err error) {
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	u, _ := url.ParseQuery(r.URL.RawQuery)

	var card model.Card

	// クエリパラメータを取得
	card.AnswerText = u["answerText"][0]
	card.QuestionText = u["questionText"][0]

	
	card.LearningLevel = 0
	card.TagId = 1
	card.CreatedUserId=2

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

// 既存のカードを一件削除
func DeleteOneCard(w http.ResponseWriter, r *http.Request) (err error) {
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// urlからidを取得
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
		fmt.Println("success to delete card id= ", id)
		w.WriteHeader(200)
	}

	return err
}

// カード情報を更新
func UpdateOneCard(w http.ResponseWriter, r *http.Request) (err error) {
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	u, _ := url.ParseQuery(r.URL.RawQuery)

	// query -> map[a:[AAA] b:[BBB] c:[CCC] d:[DDD]]

	var card model.Card

	// クエリパラメータを取得
	vars := mux.Vars(r)
	id := vars["id"]

	card.AnswerText = u["answerText"][0]
	card.QuestionText = u["questionText"][0]
	card.CardId, err = strconv.Atoi(id)
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
	} else {
		w.WriteHeader(200)
	}

	return err
}
