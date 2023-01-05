package models

type User struct {
	Base
	UserId      string `json:"user_id,omitempty"`
	Email       string `json:"email"`
	FullName    string `json:"full_name"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	Designation string `json:"designation"`
	Password    string `json:"password"`
}

// TableName gives table name of model
func (m User) TableName() string {
	return "users"
}

// ToMap convert User to map
func (m User) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"user_id":      m.UserId,
		"email":        m.Email,
		"full_name":    m.FullName,
		"phone_number": m.PhoneNumber,
		"address":      m.Address,
		"designation":  m.Designation,
	}
}
