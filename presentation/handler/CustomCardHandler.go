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

	fmt.Println("HandlerCardRequest")
	var err error
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
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	tokenString := r.Header.Get("Authorization")
	token := tokenString[7:]
	user, err := auth.GetUserFromBearerToken(token)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("user info from token userName =", user.UserName, "id =", user.UserId)
	cards, err := repository.GetAllCardsByUserId_DB(user.UserId)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("result=", cards)
	output, err := json.MarshalIndent(&cards, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}

	w.Write(output)
	return
}
