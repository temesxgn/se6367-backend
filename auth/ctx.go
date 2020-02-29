package auth

import (
	"context"
	"github.com/temesxgn/se6367-backend/auth/models"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

type contextKey string

// possible context key values
const (
	UserIDCtxKey        = contextKey("x-hasura-user-id")
	UserRoleCtxKey      = contextKey("x-hasura-role")
	AdminSecretCtxKey   = contextKey("x-hasura-admin-secret")
	AppCtxKey           = contextKey("x-hasura-app")
	APIKey              = contextKey("x-api-key")
	APISecret           = contextKey("x-api-secret")
	AuthorizationCtxKey = contextKey("authorization")
	UserCtxKey          = contextKey("user")
)

// String - string representation of context key
func (c contextKey) String() string {
	return string(c)
}

// SetValuesFromHeaders - Sets the header values to the context
func SetValuesFromHeaders(req *http.Request) context.Context {
	reqID := req.Header.Get(echo.HeaderXRequestID)
	if reqID == "" {
		reqID = uuid.New().String()
	}

	token := req.Header.Get(AuthorizationCtxKey.String())
	ctx := context.WithValue(req.Context(), AdminSecretCtxKey, req.Header.Get(AdminSecretCtxKey.String()))
	ctx = context.WithValue(ctx, APIKey, req.Header.Get(APIKey.String()))
	ctx = context.WithValue(ctx, APISecret, req.Header.Get(APISecret.String()))
	ctx = context.WithValue(ctx, AuthorizationCtxKey, token)
	ctx = context.WithValue(ctx, echo.HeaderXRequestID, reqID)
	// ctx = context.WithValue(ctx, UserCtxKey, GetUserFromToken(token))

	return ctx
}

// GetUser - Retrieve user from context
func GetUser(ctx context.Context) *models.User {
	usr := ctx.Value(UserCtxKey)

	if user, ok := usr.(*models.User); ok {
		return user
	}

	return &models.User{}
}
