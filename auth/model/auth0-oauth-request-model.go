package model

type Auth0ClientCredentialsTokenRequest struct {
	Audience     string `json:"audience"`
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type Auth0ClientCredentialsTokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func NewAuth0TokenRequest(audience, id, secret string) *Auth0ClientCredentialsTokenRequest {
	return &Auth0ClientCredentialsTokenRequest{
		Audience:     audience,
		GrantType:    "client_credentials",
		ClientID:     id,
		ClientSecret: secret,
	}
}
