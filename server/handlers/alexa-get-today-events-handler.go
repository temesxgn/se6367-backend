package handlers

import (
	"context"
	"fmt"
	"github.com/arienmalec/alexa-go"
	ala "github.com/temesxgn/se6367-backend/alexa"
	"github.com/temesxgn/se6367-backend/auth"
	"github.com/temesxgn/se6367-backend/config"
	"github.com/temesxgn/se6367-backend/hasura"
	"github.com/temesxgn/se6367-backend/hasura/models"
	"gopkg.in/auth0.v3"
)

// GetMyEventsForTodayIntent -
func GetMyEventsForTodayIntent(user *auth.User) (alexa.Response, error) {
	var builder ala.SSMLBuilder
	service := hasura.NewService(config.GetHasuraEndpoint())
	events, _ := service.GetEvents(context.Background(), &models.EventFilterParams{
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
