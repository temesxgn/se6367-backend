package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func setupRoutes(e *echo.Echo) {
	// e.GET("/", handler.PlaygroundHandler)
	// e.POST("/query", handler.GraphqlHandler)
	// e.POST("/stripe-webhooks", handler.StripeWebhookHandler)
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "hello")
	})
}

func setupMiddleWare(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.GzipWithConfig(middleware.DefaultGzipConfig))
	// e.Use(auth.WhiteList)
}

// New - creates a new instance of the server
func New() *echo.Echo {
	e := echo.New()
	setupMiddleWare(e)
	setupRoutes(e)
	return e
}
