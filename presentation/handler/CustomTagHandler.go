package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/sakana7392/AnkiCard_server/application/auth"
	"github.com/sakana7392/AnkiCard_server/domain/model"
	"github.com/sakana7392/AnkiCard_server/infra/repository"
)

func HandleCustomTagRequest(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Handle Custom Card Request")

	var err error
	setting(w, r)
	switch r.Method {
	case "GET":
		err = GetAllTagsByUserId(w, r)
	case "POST":
		err = CreateNewTag(w, r)
		// case "PUT":
		// 	err = UpdateOneCard(w, r)
		// case "DELETE":
		// 	err = DeleteOneCard(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	return
}

// カードを1件取得
func GetAllTagsByUserId(w http.ResponseWriter, r *http.Request) (err error) {

	// Bearer tokenからユーザ情報を取得
	tokenString := r.Header.Get("Authorization")
	token := tokenString[7:]
	user, err := auth.GetUserFromBearerToken(token)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("user info from token userName =", user.UserName, "id =", user.UserId)

	//	ユーザが作成したカードを取得
	tags, err := repository.GetAllTagsByUserId_DB(user.UserId)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("result=", tags)
	//	jsonにエンコードする
	output, err := json.MarshalIndent(&tags, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}

	w.Write(output)
	w.WriteHeader(200)
	return
}
func CreateNewTag(w http.ResponseWriter, r *http.Request) (err error) {
	// Bearer tokenからユーザ情報を取得
	tokenString := r.Header.Get("Authorization")
	token := tokenString[7:]
	user, err := auth.GetUserFromBearerToken(token)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("user info from token userName =", user.UserName, "id =", user.UserId)

	var tag model.Tag
	tag.CreatedUserId = user.UserId
	//	クエリパラメータからタグ名を取得
	u, _ := url.ParseQuery(r.URL.RawQuery)
	tag.TagName = u["tag_name"][0]

	// 同じ名前のタグが存在しないかチェック
	exist, err := repository.IfExistTagName(tag.TagName)
	if err != nil {
		log.Println(err)
		fmt.Println("failed to check new tag!")
	}
	if exist {
		fmt.Println("same tag_name is already exist!")
		w.WriteHeader(400)
		output, error := json.MarshalIndent("same tag_name is already exist!", "", "\t")
		if error != nil {
			log.Println(err)
		}
		w.Write(output)
		return
	}

	if err != repository.CreateNewTag_DB(&tag) {
		log.Println(err)
		fmt.Println("failed to create new tag!")
	}
	return
}
