package handlers

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo"
	"github.com/temesxgn/se6367-backend/auth"
	"github.com/temesxgn/se6367-backend/graph"
	"github.com/temesxgn/se6367-backend/graph/generated"
)

// GraphqlHandler - The Graphql handler
func GraphqlHandler(c echo.Context) error {
	cfg := generated.Config{Resolvers: &graph.Resolver{}}
	cfg.Directives.HasRole = hasRoleHandler(c)
	cfg.Directives.IsAuthenticated = isAuthenticatedHandler(c)

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

func hasRoleHandler(c echo.Context) func(ctx context.Context, obj interface{}, next graphql.Resolver, role auth.Role) (res interface{}, err error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver, role auth.Role) (interface{}, error) {
		ctx = auth.SetValuesFromHeaders(c.Request())
		user := auth.GetUser(ctx)
		if auth.HasAdminSecret(ctx) || user.IsValid() && user.HasRole(&role) {
			return next(ctx)
		}

		return nil, fmt.Errorf("access denied: insufficient authorization")
	}
}

func isAuthenticatedHandler(c echo.Context) func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	return func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		ctx = auth.SetValuesFromHeaders(c.Request())
		user := auth.GetUser(ctx)
		if auth.HasAdminSecret(ctx) || user.IsValid() {
			return next(ctx)
		}

		return nil, fmt.Errorf("access denied: must be authenticated")
	}
}
