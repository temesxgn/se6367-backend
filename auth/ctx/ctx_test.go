package ctx_test

import (
	"context"
	"github.com/temesxgn/se6367-backend/auth/ctx"
	"github.com/temesxgn/se6367-backend/auth/model"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/temesxgn/se6367-backend/common/util/jsonutils"
)

func TestCtx(t *testing.T) {
	testGetUser(t)
	testSetValuesFromHeaders(t)
	testContextKeyToString(t)
}

func testGetUser(t *testing.T) {
	validUser := &model.User{}
	validUserCtx := context.WithValue(context.Background(), ctx.UserCtxKey, validUser)

	nilUserCtx := context.WithValue(context.Background(), ctx.UserCtxKey, nil)
	malformedUserCtx := context.WithValue(context.Background(), ctx.UserCtxKey, "")
	tables := []struct {
		name string
		data context.Context
		want *model.User
	}{
		{"get valid user from context", validUserCtx, validUser},
		{"get nil user from context", nilUserCtx, &model.User{}},
		{"get malformed user from context", malformedUserCtx, &model.User{}},
	}

	for _, tt := range tables {
		t.Run(tt.name, func(t *testing.T) {
			got := ctx.GetUser(tt.data)
			gJSON, _ := jsonutils.Marshal(got)
			wJSON, _ := jsonutils.Marshal(tt.want)
			assert.EqualValues(t, got, tt.want, "Expected %s, Actual %s", wJSON, gJSON)
		})
	}
}

func testSetValuesFromHeaders(t *testing.T) {
	req, _ := http.NewRequest("", "", nil)
	req.Header.Add(ctx.AuthorizationCtxKey.String(), "")

	tables := []struct {
		name string
		data *http.Request
		want context.Context
	}{}

	for _, tt := range tables {
		t.Run(tt.name, func(t *testing.T) {
			got := ctx.SetValuesFromHeaders(tt.data)
			gJSON, _ := jsonutils.Marshal(got)
			wJSON, _ := jsonutils.Marshal(tt.want)
			assert.EqualValues(t, got, tt.want, "Expected %s, Actual %s", wJSON, gJSON)
		})
	}
}

func testContextKeyToString(t *testing.T) {

}
