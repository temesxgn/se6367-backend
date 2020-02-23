package auth_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/temesxgn/se6367-backend/auth"
	"github.com/temesxgn/se6367-backend/util/jsonutils"
)

func TestCtx(t *testing.T) {
	testGetUser(t)
	testSetValuesFromHeaders(t)
	testContextKeyToString(t)
}

func testGetUser(t *testing.T) {
	validUser := &auth.User{}
	validUserCtx := context.WithValue(context.Background(), auth.UserCtxKey, validUser)

	nilUserCtx := context.WithValue(context.Background(), auth.UserCtxKey, nil)
	malformedUserCtx := context.WithValue(context.Background(), auth.UserCtxKey, "")
	tables := []struct {
		name string
		data context.Context
		want *auth.User
	}{
		{"get valid user from context", validUserCtx, validUser},
		{"get nil user from context", nilUserCtx, &auth.User{}},
		{"get malformed user from context", malformedUserCtx, &auth.User{}},
	}

	for _, tt := range tables {
		t.Run(tt.name, func(t *testing.T) {
			got := auth.GetUser(tt.data)
			gJSON, _ := jsonutils.Marshal(got)
			wJSON, _ := jsonutils.Marshal(tt.want)
			assert.EqualValues(t, got, tt.want, "Expected %s, Actual %s", wJSON, gJSON)
		})
	}
}

func testSetValuesFromHeaders(t *testing.T) {
	req, _ := http.NewRequest("", "", nil)
	req.Header.Add(auth.AuthorizationCtxKey.String(), "")

	tables := []struct {
		name string
		data *http.Request
		want context.Context
	}{}

	for _, tt := range tables {
		t.Run(tt.name, func(t *testing.T) {
			got := auth.SetValuesFromHeaders(tt.data)
			gJSON, _ := jsonutils.Marshal(got)
			wJSON, _ := jsonutils.Marshal(tt.want)
			assert.EqualValues(t, got, tt.want, "Expected %s, Actual %s", wJSON, gJSON)
		})
	}
}

func testContextKeyToString(t *testing.T) {

}
