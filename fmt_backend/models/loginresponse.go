package models

type LoginResponseModel struct {
	AccessToken string `json:"access_token"`
}

func (m LoginResponseModel) ToMapR() map[string]interface{} {
	return map[string]interface{}{
		"access_token": m.AccessToken,
	}
}
