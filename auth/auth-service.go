package auth

import (
	"github.com/temesxgn/se6367-backend/auth/auth0"
	"github.com/temesxgn/se6367-backend/graph/model"
)

type Service interface {
	GetToken() (string, error)
	GetUser(userID string) (*model.Auth0Profile, error)
	CreateUser(connection, email string) error
}

func GetAuthService(aType ServiceType) (Service, error) {
	switch aType {
	case AuthZeroAuthServiceType:
		fallthrough
	default:
		return auth0.NewService()
	}
}
