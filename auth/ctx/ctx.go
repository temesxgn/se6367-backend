package ctx

import (
	"context"
	"errors"
	"github.com/temesxgn/se6367-backend/auth/model"
	"gopkg.in/square/go-jose.v2/jwt"
	"net/http"
	"strings"

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
	usr, _ := GetUserFromToken(token)
	ctx := context.WithValue(req.Context(), UserCtxKey, usr)
	ctx = context.WithValue(ctx, AdminSecretCtxKey, req.Header.Get(AdminSecretCtxKey.String()))
	ctx = context.WithValue(ctx, APIKey, req.Header.Get(APIKey.String()))
	ctx = context.WithValue(ctx, APISecret, req.Header.Get(APISecret.String()))
	ctx = context.WithValue(ctx, AuthorizationCtxKey, token)
	ctx = context.WithValue(ctx, echo.HeaderXRequestID, reqID)

	return ctx
}

// GetUser - Retrieve user from context
func GetUser(ctx context.Context) *model.User {
	usr := ctx.Value(UserCtxKey)

	if user, ok := usr.(*model.User); ok {
		return user
	}

	return &model.User{}
}

func GetUserFromToken(token string) (*model.User, error) {
	if split := strings.Split(token, " "); len(split) == 2 {
		token := split[1]
		tkn, err := jwt.ParseSigned(token)
		if err != nil {
			return nil, err
		}

		usr := &model.User{}
		if err := tkn.UnsafeClaimsWithoutVerification(usr); err != nil {
			return nil, err
		}

		return usr, nil
	}

	return nil, errors.New("malformed authentication request")
}
