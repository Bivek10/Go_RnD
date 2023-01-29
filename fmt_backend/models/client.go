package models

type Clients struct {
	Base
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	Address   string `json:"address"`
	Password  string `json:"password"`
}

func (m Clients) TableName() string {
	return "clients"
}

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ClientRequestResponse struct {
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Email        string `json:"email"`
	Address      string `json:"address"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
