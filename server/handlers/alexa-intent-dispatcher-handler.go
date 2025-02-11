package handlers

import (
	"github.com/temesxgn/se6367-backend/alexa/handlers"
	"github.com/temesxgn/se6367-backend/auth/ctx"
	"github.com/temesxgn/se6367-backend/auth/model"
	"net/http"

	"github.com/arienmalec/alexa-go"
	"github.com/labstack/echo"
	ala "github.com/temesxgn/se6367-backend/alexa"
)

// AlexaIntentHandler - endpoint to handle requests from alexa
func AlexaIntentHandler(c echo.Context) error {
	request := ctx.GetAlexRequestFromContext(c)
	usr := ctx.GetUserFromContext(c)

	res, err := intentDispatcher(request, usr)
	if err != nil {
		return c.JSON(http.StatusOK, ala.NewSSMLResponse("Intent Error", err.Error()))
	}

	return c.JSON(http.StatusOK, res)
}

// triggers the event matching the incoming intent request
func intentDispatcher(request *alexa.Request, usr *model.User) (alexa.Response, error) {
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
