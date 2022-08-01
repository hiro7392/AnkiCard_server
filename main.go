package main

import (
	"net/http"
	"github.com/sakana7392/AnkiCard_server/handler"
)

func main() {
	var localURL string = "127.0.0.1:8080"
	URL := localURL

	server := http.Server{
		Addr: URL,
	}
	//カードのCRUD処理
	http.HandleFunc("/card/", HandleCardRequest)

	server.ListenAndServe()
}
