package operations

import (
	"log"

	"bivek.com/fmt_api/helper"
	"bivek.com/fmt_api/interfaces"
)

func Login(email string, pass string) {
	db := helper.ConnectDB()
	user := interfaces.User{}
	// get first record order by primary key
	data := db.First(&user)

	log.Printf("%v", data)

	// if db.Where("Email=?", email).First(&user) == nil {
	// 	print("Email not found")
	// }

}
