package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID          string         `json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Email       string         `json:"email"`
	FullName    string         `json:"full_name"`
	PhoneNumber string         `json:"phone_number"`
	Password    string         `json:"password"`
}

// TableName gives table name of model
func (m User) TableName() string {
	return "users"
}

// ToMap convert User to map
func (m User) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"email":        m.Email,
		"full_name":    m.FullName,
		"phone_number": m.PhoneNumber,
	}
}
