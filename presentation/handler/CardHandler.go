package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sakana7392/AnkiCard_server/application/auth"
	"github.com/sakana7392/AnkiCard_server/domain/model"
	"github.com/sakana7392/AnkiCard_server/infra/repository"
)

func setting(w http.ResponseWriter, r *http.Request) {
	// cors用の設定
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	//プリフライトリクエストへの応答
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
}
func HandleCardRequest(w http.ResponseWriter, r *http.Request) {

	fmt.Println("HandlerCardRequest")
	setting(w, r)
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
	u, _ := url.ParseQuery(r.URL.RawQuery)

	var card model.Card

	// クエリパラメータを取得
	card.AnswerText = u["answerText"][0]
	card.QuestionText = u["questionText"][0]
	TagIdStr := u["tagId"][0]
	card.TagId, err = strconv.Atoi(TagIdStr)
	if err != nil {
		log.Println(err)
		return
	}
	// Bearer tokenからユーザ情報を取得
	tokenString := r.Header.Get("Authorization")
	token := tokenString[7:]
	user, err := auth.GetUserFromBearerToken(token)
	if err != nil {
		log.Println(err)
	}

	card.CreatedUserId = user.UserId

	// tagNameをtagIdから取得
	tag, err := repository.GetOneTag_DB(card.TagId)
	if err != nil {
		log.Println(err)
	}
	card.TagName = tag.TagName
	card.LearningLevel = 0
	fmt.Println("card.TagName =", card.TagName)

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

	// urlからidを取得

	vars := mux.Vars(r)
	id := vars["id"]

	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		return
	}
	err = repository.DeleteOneCard_DB(idInt)
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
