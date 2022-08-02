package auth

import (
	"net/http"
	"os"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/form3tech-oss/jwt-go"
)

// GetTokenHandler get token
var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

    // headerのセット
	token := jwt.New(jwt.SigningMethodHS256)

    // claimsのセット
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["sub"] = "54546557354"
	claims["name"] = "taro"
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

    // 電子署名
	tokenString, _ := token.SignedString([]byte(os.Getenv("SIGNINGKEY")))

    // JWTを返却
	w.Write([]byte(tokenString))
})

// JwtMiddleware check token
var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SIGNINGKEY")), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})