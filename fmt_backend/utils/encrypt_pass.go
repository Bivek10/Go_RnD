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

func DecryptPassword(hasedpass []byte, pass []byte) bool {
	err := bcrypt.CompareHashAndPassword(hasedpass, pass)
	print(err)
	if err != nil {
		return false
	}
	return true
}
