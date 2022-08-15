package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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
		err = GetAllCardsByUserId(w, r)
		// case "POST":
		// 	err = CreateNewCard(w, r)
		// case "PUT":
		// 	err = UpdateOneCard(w, r)
		// case "DELETE":
		// 	err = DeleteOneCard(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
