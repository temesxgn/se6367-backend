package models

import "fmt"

type Auth0ClientCredentialsTokenRequest struct {
	Audience     string `json:"audience"`
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type Auth0ClientCredentialsTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

func NewAuth0TokenRequest(domain, id, secret string) *Auth0ClientCredentialsTokenRequest {
	return &Auth0ClientCredentialsTokenRequest{
		Audience:     fmt.Sprintf("https://%v/oauth/token", domain),
		GrantType:    "client_credentials",
		ClientID:     id,
		ClientSecret: secret,
	}
}
