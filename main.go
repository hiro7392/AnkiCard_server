package main

import (
	"net/http"
	//"strconv"
	"github.com/sakana7392/AnkiCard_server/handler"
)

func handleCardRequest(w http.ResponseWriter, r *http.Request) {

	var err error
	switch r.Method {
	case "GET":
		err = getOneCard(w, r)
	case "POST":
		err = createNewCard(w, r)
	case "PUT":
		err = updateOneCard(w, r)
	case "DELETE":
		err = deleteOneCard(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	return
}

func main() {
	var localURL string = "127.0.0.1:8080"
	URL := localURL

	server := http.Server{
		Addr: URL,
	}
	//カードのCRUD処理
	http.HandleFunc("/card/", handleCardRequest)

	server.ListenAndServe()
}
