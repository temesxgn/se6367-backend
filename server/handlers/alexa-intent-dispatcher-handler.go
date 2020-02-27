package handlers

import (
	"net/http"

	"github.com/arienmalec/alexa-go"
	"github.com/labstack/echo"
	ala "github.com/temesxgn/se6367-backend/alexa"
	"github.com/temesxgn/se6367-backend/auth"
)

// AlexaIntentHandler - endpoint to handle requests from alexa
func AlexaIntentHandler(c echo.Context) error {
	request := c.Get("request").(*alexa.Request)
	usr := c.Get("user").(*auth.User)

	res, err := IntentDispatcher(request, usr)
	if err != nil {
		return c.JSON(http.StatusOK, ala.NewSSMLResponse("Intent Error", err.Error()))
	}

	return c.JSON(http.StatusOK, res)
}

// IntentDispatcher -
func IntentDispatcher(request *alexa.Request, usr *auth.User) (alexa.Response, error) {
	switch request.Body.Intent.Name {
	case ala.GetMyEventsForTodayIntentType.String():
		return GetMyEventsForTodayIntent(usr)
	default:
		return HandleHelpIntent()
	}
}
