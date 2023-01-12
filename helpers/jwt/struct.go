package jwt

import "simple-api/helpers"

type (
	Token struct {
		AccessToken      string      `json:"access_token"`
		ExpiresIn        interface{} `json:"expires_in"`
		RefreshExpiresIn interface{} `json:"refresh_expires_in"`
		RefreshToken     string      `json:"refresh_token"`
		TokenType        string      `json:"token_type"`
	}
	TokenDetail struct {
		Id        *helpers.UUID `json:"id"`
		Firstname string        `json:"firstname"`
		Lastname  string        `json:"lastname"`
	}
	RefreshToken struct {
		RefreshExpiresIn interface{} `json:"refresh_expires_in"`
		RefreshToken     string      `json:"refresh_token"`
	}
)
