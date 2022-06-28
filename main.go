package main

import (
	"fmt"
	"net/http"
	"net/url"
	//"strconv"
)

func hello(w http.ResponseWriter, r *http.Request) (err error) {
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Printf("first called\n")
	w.WriteHeader(200)
	return
}

//	カードを全権取得
func getCardsOfUser(w http.ResponseWriter, r *http.Request) (err error) {
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	fmt.Println("we will return cards all of user")
	w.WriteHeader(200)
	return
}

//	カードを新規作成
func createNewCard(w http.ResponseWriter, r *http.Request) (err error) {
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	u, _ := url.ParseQuery(r.URL.RawQuery)
	fmt.Println(u)

	// query -> map[a:[AAA] b:[BBB] c:[CCC] d:[DDD]]

	var card Card
	card.AnswerText = u["answerText"][0]
	card.QuestionText = u["questionText"][0]
	card.LearningLevel = 0
	card.TagId = 1

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
		err = hello(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("first called\n")
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
