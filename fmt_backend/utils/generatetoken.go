package utils

import (
	"time"

	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(email string) (string, error, string, error) {
	var env infrastructure.Env
	var mySigningKey = []byte(env.JWTSecretKey)
	var err error
	clams := jwt.MapClaims{}
	clams["email"] = email
	clams["exp"] = time.Now().Add(time.Minute * 30).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clams)
	accessToken, err := token.SignedString(mySigningKey)

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["email"] = email
	rtClaims["exp"] = time.Now().Add(time.Hour * 48).Unix()

	refreshtoken, refresherr := refreshToken.SignedString([]byte(env.JWRTSecretKey))

	return accessToken, err, refreshtoken, refresherr
}
