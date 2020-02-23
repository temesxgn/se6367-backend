package handlers

import (
	"fmt"
	"net/http"

	"github.com/arienmalec/alexa-go"
	"github.com/labstack/echo"
	ala "github.com/temesxgn/se6367-backend/alexa"
	"github.com/temesxgn/se6367-backend/util/jsonutils"
)

func AlexaIntentHandler(c echo.Context) error {
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

// IntentDispatcher -
func IntentDispatcher(request *alexa.Request) (alexa.Response, error) {
	js, _ := jsonutils.Marshal(request)
	fmt.Println("Request /n" + js)
	switch request.Body.Intent.Name {
	case ala.GetMyEventsForTodayIntentType.String():
		return GetMyEventsForTodayIntent(request)
	default:
		return HandleHelpIntent(request)
	}
}
