package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sakana7392/AnkiCard_server/application/auth"
	"github.com/sakana7392/AnkiCard_server/infra/repository"
)

func HandleCustomCardRequest(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Handle Custom Card Request")

	var err error
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
	switch r.Method {
	case "GET":
		// ユーザIDからそのユーザが作成したカードを取得
		err = GetAllCardsByUserId(w, r)
		// case "POST":
		// 	err = CreateNewCard(w, r)
	case "PUT":
		// カードのレベルを更新
		err = UpdateOneCardLevel(w, r)
		// case "DELETE":
		// 	err = DeleteOneCard(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.WriteHeader(200)
	return
}

// カードを1件取得
func GetAllCardsByUserId(w http.ResponseWriter, r *http.Request) (err error) {

	//プリフライトリクエストへの応答
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Bearer tokenからユーザ情報を取得
	tokenString := r.Header.Get("Authorization")
	token := tokenString[7:]
	user, err := auth.GetUserFromBearerToken(token)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("user info from token userName =", user.UserName, "id =", user.UserId)

	//	ユーザが作成したカードを取得
	cards, err := repository.GetAllCardsByUserId_DB(user.UserId)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("result=", cards)
	//	jsonにエンコードする
	output, err := json.MarshalIndent(&cards, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}

	w.Write(output)
	return
}

// カード情報を更新
func UpdateOneCardLevel(w http.ResponseWriter, r *http.Request) (err error) {

	// クエリパラメータ(levelとカードID)を取得
	vars := mux.Vars(r)
	addLevelStr := r.FormValue("level")
	println("addLevelStr=", addLevelStr)
	IdStr := vars["id"]

	addLevelInt, err := strconv.Atoi(addLevelStr)
	if err != nil {
		log.Println(err)
		return
	}
	IdInt, err := strconv.Atoi(IdStr)
	if err != nil {
		log.Println(err)
		return
	}
	//	カードのレベルを更新
	card, err := repository.UpdateOneCardLevel_DB(addLevelInt, IdInt)
	if err != nil {
		fmt.Println("failed to update card")
		fmt.Println(err)
		w.WriteHeader(500)
		return
	}
	//	jsonにエンコードする
	output, err := json.MarshalIndent(&card, "", "\t")
	if err != nil {
		log.Println(err)
	}
	w.Write(output)
	return err
}
