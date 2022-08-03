package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sakana7392/AnkiCard_server/auth"
	"github.com/sakana7392/AnkiCard_server/handler"
)

type post struct {
	Title string `json:"title"`
	Tag   string `json:"tag"`
	URL   string `json:"url"`
}

func main() {
	// var localURL string = "127.0.0.1:8080"
	// URL := localURL

	// server := http.Server{
	// 	Addr: URL,
	// }
	// //カードのCRUD処理
	// http.HandleFunc("/card/", handler.HandleCardRequest)

	// //ユーザ情報の編集・参照
	// http.HandleFunc("/user/", handler.HandleUserRequest)

	// //ユーザのログイン処理
	// http.HandleFunc("/auth", auth.GetTokenHandler)

	// //ユーザのログアウト処理
	// //http.HandleFunc("/private/user/", auth.JwtMiddleware.Handler(handler.privateCardFunc))

	// server.ListenAndServe()

	r := mux.NewRouter()

	// 認証
	r.Handle("/auth", auth.GetTokenHandler)
	// 認証テスト用
	//http.HandlerFunc(handler.HandleCardRequest)

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
