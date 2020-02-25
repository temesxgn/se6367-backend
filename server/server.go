package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/temesxgn/se6367-backend/server/handlers"
)

func setupRoutes(e *echo.Echo) {
	e.POST("/alexa", handlers.AlexaIntentHandler)
	//e.POST("/insert", handlers.InsertEventHandler)
}

func setupMiddleWare(e *echo.Echo) {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.GzipWithConfig(middleware.DefaultGzipConfig))
}

// New - creates a new instance of the server
func New() *echo.Echo {
	e := echo.New()
	setupMiddleWare(e)
	setupRoutes(e)
	return e
}
