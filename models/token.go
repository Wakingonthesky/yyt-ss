package models

type Token struct {
	AccessToken string `json:"access_token" form:"access_token"`
}
