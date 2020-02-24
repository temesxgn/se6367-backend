package auth

import "gopkg.in/auth0.v3/management"

type Service interface {
	GetToken() (string, error)
	GetUser(userID string) (*management.User, error)
	CreateUser(connection, email string) error
}

func NewService() Service {
	return NewDefaultService()
}
