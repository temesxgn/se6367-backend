package handlers

import (
	"net/http"

	"github.com/arienmalec/alexa-go"
	"github.com/labstack/echo"
)

func InsertEventHandler(c echo.Context) error {
	u := new(alexa.Request)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	res, err := IntentDispatcher(u)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
