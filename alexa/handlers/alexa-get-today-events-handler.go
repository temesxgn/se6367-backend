package handlers

import (
	"context"
	"fmt"
	"github.com/arienmalec/alexa-go"
	ala "github.com/temesxgn/se6367-backend/alexa"
	"github.com/temesxgn/se6367-backend/auth/model"
	models2 "github.com/temesxgn/se6367-backend/common/models"
	"github.com/temesxgn/se6367-backend/event"
	"gopkg.in/auth0.v3"
)

// GetMyEventsForTodayIntent -
func GetMyEventsForTodayIntentHandler(user *model.User) (alexa.Response, error) {
	var builder ala.SSMLBuilder
	service, _ := event.GetEventService(event.HasuraEventServiceType)
	events, _ := service.GetEvents(context.Background(), &models2.EventFilterParams{
		UserID: auth0.String(user.Sub),
	})

	if len(events) == 0 {
		builder.Say("You have no events for today.")
		builder.Pause("300")
		builder.Say("To create an event, say Alexa, events manager create event")
	} else {
		builder.Say("Here are your events for today")
		builder.Pause("500")
		for _, event := range events {
			builder.Say(fmt.Sprintf("%s", event.Title))
			builder.Pause("1000")
		}
	}

	return ala.NewSSMLResponse("My Events Today", builder.Build()), nil
}
