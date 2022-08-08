package auth

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/form3tech-oss/jwt-go"
	"github.com/sakana7392/AnkiCard_server/application/service"
)

// GetTokenHandler get token
var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// headerのセット
	token := jwt.New(jwt.SigningMethodHS256)

	//	クエリパラメータからemailとpasswordを取得
	u, Er := url.ParseQuery(r.URL.RawQuery)
	if Er !=nil{
		log.Println(Er)
		return
	}else{
		fmt.Println("email:", u.Get("email"))
		fmt.Println("password:", u.Get("password"))
	}
	receivedEmail := u.Get("email")
	receivedPassword := u.Get("password")
	

	//	emailとpasswordが存在するかチェック
	if !service.CheckEmailAndPassword(receivedEmail, receivedPassword) {
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		fmt.Println("emailとユーザが存在します")
	}

	// claimsのセット
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = u["email"][0]
	claims["password"] = u["password"][0]

	// 電子署名
	tokenString, _ := token.SignedString([]byte(os.Getenv("SIGNINGKEY")))

	// JWTを返却
	w.Write([]byte(tokenString))
	
	// サーバだけが知り得るSecretでこれをParseする
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNINGKEY")), nil
	})
	if err != nil {
		fmt.Println("jwt.Parse error ",err)
	}
})

// JwtMiddleware check token
var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNINGKEY")), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})
