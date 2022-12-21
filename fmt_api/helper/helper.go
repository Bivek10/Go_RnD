package helper
import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
  )

func connnetDB() *gorm.DB{
	db, err:=gorm.Open()
}