package interfaces

import "gorm.io/gorm"

type User struct{
	gorm.Model
	Username string 
	Phonenumber string
	Email string
	Password string
	FarmerType string

}