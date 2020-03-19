package handlers

import (
	"bytes"
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo"
	authCtx "github.com/temesxgn/se6367-backend/auth/ctx"
	"github.com/temesxgn/se6367-backend/auth/middleware"
	"github.com/temesxgn/se6367-backend/auth/model"
	"github.com/temesxgn/se6367-backend/graph"
	"github.com/temesxgn/se6367-backend/graph/generated"
)

// GraphqlHandler - The Graphql handler
func GraphqlHandler(c echo.Context) error {
	cfg := generated.Config{Resolvers: &graph.Resolver{}}
	cfg.Directives.HasRole = hasRoleHandler(c)
	cfg.Directives.IsAuthenticated = isAuthenticatedHandler(c)
	cfg.Directives.HasIntegration = hasIntegration(c)

	h := handler.NewDefaultServer(generated.NewExecutableSchema(cfg))
	req := c.Request()
	res := c.Response()
	h.ServeHTTP(res, req)

	return nil
}

// PlaygroundHandler - The Playground handler
func PlaygroundHandler(c echo.Context) error {
	h :=  playground.Handler("GraphQL", "/query")
	req := c.Request()
	res := c.Response()
	h.ServeHTTP(res, req)

	return nil
}

func hasRoleHandler(c echo.Context) func(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (res interface{}, err error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (interface{}, error) {
		ctx = authCtx.SetValuesFromHeaders(c.Request())
		user := authCtx.GetUser(ctx)
		if middleware.HasAdminSecret(ctx) || user.IsValid() && user.HasRole(&role) {
			return next(ctx)
		}

		return nil, fmt.Errorf("access denied: insufficient authorization")
	}
}

func isAuthenticatedHandler(c echo.Context) func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		ctx = authCtx.SetValuesFromHeaders(c.Request())
		user := authCtx.GetUser(ctx)
		if middleware.HasAdminSecret(ctx) || user.IsValid() {
			return next(ctx)
		}

		return nil, fmt.Errorf("access denied: must be authenticated")
	}
}

func hasIntegration(c echo.Context) func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		buf := new(bytes.Buffer)
		buf.ReadFrom(c.Request().Body)
		fmt.Println(fmt.Sprintf("OBJECT: %v", buf.String()))
		return next(ctx)
	}
}
