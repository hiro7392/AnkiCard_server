package main

import (
	"fmt"
	"net/http"
	//"strconv"
)

//	カードを全権取得
func getCardsOfUser(w http.ResponseWriter, r *http.Request) (err error) {
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Println("we will return cards all of user")
	w.WriteHeader(200)
	return
}

func handleRequest(w http.ResponseWriter, r *http.Request) {

	var err error
	switch r.Method {
	case "GET":
		err = getCardsOfUser(w, r)
	case "POST":
		err = createNewCard(w, r)
	case "PUT":
		err = hello(w, r)
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
	http.HandleFunc("/", handleRequest)

	server.ListenAndServe()
}
