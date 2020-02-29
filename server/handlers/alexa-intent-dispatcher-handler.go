package handlers

import (
	"github.com/temesxgn/se6367-backend/alexa/handlers"
	"github.com/temesxgn/se6367-backend/auth/models"
	"net/http"

	"github.com/arienmalec/alexa-go"
	"github.com/labstack/echo"
	ala "github.com/temesxgn/se6367-backend/alexa"
)

// AlexaIntentHandler - endpoint to handle requests from alexa
func AlexaIntentHandler(c echo.Context) error {
	request := c.Get("request").(*alexa.Request)
	usr := c.Get("user").(*models.User)

	res, err := intentDispatcher(request, usr)
	if err != nil {
		return c.JSON(http.StatusOK, ala.NewSSMLResponse("Intent Error", err.Error()))
	}

	return c.JSON(http.StatusOK, res)
}

// triggers the event matching the incoming intent request
func intentDispatcher(request *alexa.Request, usr *models.User) (alexa.Response, error) {
	switch request.Body.Intent.Name {
	case ala.CreateEventIntentType.String():
		return handlers.CreateEventIntentHandler(request, usr)
	case ala.GetMyEventsForTodayIntentType.String():
		return handlers.GetMyEventsForTodayIntentHandler(usr)
	case ala.GetEventsForDayIntentType.String():
		return handlers.GetMyEventsForDayIntentHandler(request, usr)
	case ala.DeleteEventIntentType.String():
		return handlers.DeleteEventIntentHandler(request, usr)
	default:
		return handlers.HandleHelpIntentHandler()
	}
}
