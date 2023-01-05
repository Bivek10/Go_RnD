package models

type UserLoginModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (m User) ToMapD() map[string]interface{} {
	return map[string]interface{}{
		"email":    m.Email,
		"password": m.Password,
	}
}
