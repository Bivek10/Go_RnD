package operations

import (
	"bivek.com/fmt_api/helper"
	"bivek.com/fmt_api/interfaces"
)

func CreateAccount() {
	db := helper.ConnectDB()
	users := [1]interfaces.User{
		{Username: "Bivek Karki", Phonenumber: "9800965652", Email: "bibekkarki662@gmail.com", Password: "test1234", FarmerType: "Commerical"},
	}

	for i := 0; i < len(users); i++ {
		encryptPassword := helper.EncryptPassword([]byte(users[i].Password))

		user := &interfaces.User{Username: users[i].Username, Phonenumber: users[i].Phonenumber, Password: encryptPassword, Email: users[i].Email, FarmerType: users[i].FarmerType}

		db.Create(&user)

	}

	db.Commit()
}

func Migrate() {
	User := &interfaces.User{}
	db := helper.ConnectDB()

	db.AutoMigrate(&User)

	CreateAccount()

}
