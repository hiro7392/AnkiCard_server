package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sakana7392/AnkiCard_server/application/auth"

	"github.com/sakana7392/AnkiCard_server/application/service"
	"github.com/sakana7392/AnkiCard_server/presentation/handler"
)

type post struct {
	Title string `json:"title"`
	Tag   string `json:"tag"`
	URL   string `json:"url"`
}

func main() {

	r := mux.NewRouter()
	// テストデータの挿入
	service.InsertTestUserData()
	service.InsertTestTagData()
	service.InsertTestCardData()
	// 認証
	r.Handle("/auth", auth.GetTokenHandler)

	// カードのCRUD処理

	// 認証あり
	r.Handle("/private/card/{id}", auth.JwtMiddleware.Handler(cardAuth))
	// 認証なし
	r.Handle("/card/{id}", cardwithoutAuth)

	//サーバー起動
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("ListenAndServe:", err)
	}

}

var cardAuth = http.HandlerFunc(handler.HandleCardRequest)
var cardwithoutAuth = http.HandlerFunc(handler.HandleCardRequest)

// var private = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 	post := &post{
// 		Title: "VGolangとGoogle Cloud Vision APIで画像から文字認識するCLIを速攻でつくる",
// 		Tag:   "Go",
// 		URL:   "https://qiita.com/po3rin/items/bf439424e38757c1e69b",
// 	}
// 	json.NewEncoder(w).Encode(post)
// })
