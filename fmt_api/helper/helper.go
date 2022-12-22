package helper

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", "root", "", "localhost", "3306", "fmt_db")

	db, err := gorm.Open(mysql.Open(url))

	HandleError(err)
	fmt.Println(db)
	fmt.Println("Success!")
	return db
}

func HandleError(err error) {
	if err != nil {
		panic(err.Error())
	}
}
