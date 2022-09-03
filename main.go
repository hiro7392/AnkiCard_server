package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sakana7392/AnkiCard_server/application/auth"

	"github.com/sakana7392/AnkiCard_server/application/service"
	"github.com/sakana7392/AnkiCard_server/presentation/handler"
)

func insertTestData() {
	service.InsertTestUserData()
	service.InsertTestTagData()
	service.InsertTestCardData()
}
func main() {

	r := mux.NewRouter()
	// テストデータの挿入
	//insertTestData()

	// JWTで認証。Bearer Tokenを発行する
	r.Handle("/auth", auth.GetTokenHandler)

	// カードのCRUD処理
	r.Handle("/card/{id}", auth.JwtMiddleware.Handler(cardAuth))

	// ユーザが作成したカードを取得
	r.Handle("/private/card/{id}", auth.JwtMiddleware.Handler(CustomCardRequest))

	// ユーザが作成したカードのレベルを更新
	r.Path("/private/card/{id}").Queries("level", "{level}").HandlerFunc(CustomCardRequest)

	// ユーザが作成したタグを全て取得
	r.Handle("/private/tag/", auth.JwtMiddleware.Handler(CustomTagRequest))

	//サーバー起動
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

var cardAuth = http.HandlerFunc(handler.HandleCardRequest)
var CustomCardRequest = http.HandlerFunc(handler.HandleCustomCardRequest)
var CustomTagRequest = http.HandlerFunc(handler.HandleCustomTagRequest)
var CardCreateRequest = http.HandlerFunc(handler.HandleCardRequest)
