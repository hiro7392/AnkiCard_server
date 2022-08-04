package auth

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/sakana7392/AnkiCard_server/application/service"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/form3tech-oss/jwt-go"
)

// GetTokenHandler get token
var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// headerのセット
	token := jwt.New(jwt.SigningMethodHS256)

	//	クエリパラメータからemailとpasswordを取得
	u, _ := url.ParseQuery(r.URL.RawQuery)
	receivedEmail := u.Get("email")
	receivedPassword := u.Get("password")

	//	emailとpasswordが存在するかチェック
	if !service.CheckEmailAndPassword(receivedEmail, receivedPassword) {
		w.WriteHeader(http.StatusBadRequest)
		return
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
	fmt.Println("token[email]=", token.Claims.(jwt.MapClaims)["email"]) //emailを表示
	// Parseメソッドを使うと、Claimsはmapとして得られる
	log.Println(token.Claims, err)
})

// JwtMiddleware check token
var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNINGKEY")), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})