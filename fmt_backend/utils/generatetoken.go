package utils

import (
	"time"

	"github.com/bivek/fmt_backend/infrastructure"
	"github.com/bivek/fmt_backend/models"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(email string) (string, error, string, error) {
	var env infrastructure.Env
	var mySigningKey = []byte(env.JWTSecretKey)
	var err error
	clams := jwt.MapClaims{}
	clams["email"] = email
	clams["exp"] = time.Now().Add(time.Minute * 15).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, clams)
	accessToken, err := token.SignedString(mySigningKey)

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["email"] = email
	rtClaims["exp"] = time.Now().Add(time.Hour * 48).Unix()
	//refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshtoken, refresherr := refreshToken.SignedString([]byte(env.JWRTSecretKey))

	return accessToken, err, refreshtoken, refresherr
}

func GenerateJWTV(email string) (string, error, string, error) {
	var env infrastructure.Env
	clams := models.SignedDetails{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}
	refreshClams :=
		models.SignedDetails{
			Email: email,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Minute * 48).Unix(),
			},
		}
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, clams).SignedString([]byte(env.JWTSecretKey))
	refreshtoken, refresherr := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClams).SignedString([]byte(env.JWRTSecretKey))
	return accessToken, err, refreshtoken, refresherr
}
