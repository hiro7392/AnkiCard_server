package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sakana7392/AnkiCard_server/application/auth"

	"github.com/sakana7392/AnkiCard_server/application/service"
	"github.com/sakana7392/AnkiCard_server/presentation/handler"
)

func insertTestDate() {
	service.InsertTestUserData()
	service.InsertTestTagData()
	service.InsertTestCardData()
}
func main() {

	r := mux.NewRouter()
	// テストデータの挿入
	//insertTestDate()

	// JWTで認証。Bearer Tokenを発行する
	r.Handle("/auth", auth.GetTokenHandler)

	// カードのCRUD処理
	r.Handle("/card/{id}", auth.JwtMiddleware.Handler(cardAuth))

	//サーバー起動
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

var cardAuth = http.HandlerFunc(handler.HandleCardRequest)
