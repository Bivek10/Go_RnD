package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(pass []byte) string {
	value, err := bcrypt.GenerateFromPassword(pass, bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	return string(value)
}
