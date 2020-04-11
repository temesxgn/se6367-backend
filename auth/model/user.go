package model

import (
	"strings"
	"time"

	"github.com/temesxgn/se6367-backend/config"
)

// User - model based off of Auth0 JWT with Hasura Claims
type User struct {
	Iss    string                 `json:"iss"`
	Sub    string                 `json:"sub"`
	Aud    interface{}            `json:"aud"`
	Iat    int                    `json:"iat"`
	Exp    int                    `json:"exp"`
	Azp    string                 `json:"azp"`
	Scope  string                 `json:"scope"`
	Claims HTTPSHasuraIoJwtClaims `json:"https://hasura.io/jwt/claims"`
}

// HasRole - check if user has the given role
func (u *User) HasRole(role *Role) bool {
	for _, r := range u.Claims.XHasuraAllowedRoles {
		if strings.EqualFold(role.String(), r) {
			return true
		}
	}

	return false
}

// HasExpired - check if token has expired
func (u *User) HasExpired() bool {
	exp := time.Unix(int64(u.Exp), 0)
	return time.Now().Unix() >= exp.Unix()
}

// UserID - returns the user id
func (u *User) UserID() string {
	return u.Claims.XHasuraUserEmail
}

// UserID - returns the user id
func (u *User) UserEmail() string {
	return u.Claims.XHasuraUserID
}

// IsValid - validates the jwt token
func (u *User) IsValid() bool {
	return strings.EqualFold(u.Iss, config.GetAuth0Domain()) && !u.HasExpired()
}

// HTTPSHasuraIoJwtClaims - Hasura specific claims
type HTTPSHasuraIoJwtClaims struct {
	XHasuraDefaultRole  string   `json:"x-hasura-default-role"`
	XHasuraAllowedRoles []string `json:"x-hasura-allowed-roles"`
	XHasuraUserID       string   `json:"x-hasura-user-id"`
	XHasuraUserEmail    string   `json:"x-hasura-user-email"`
}
