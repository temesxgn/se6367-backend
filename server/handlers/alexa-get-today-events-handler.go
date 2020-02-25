package handlers

import (
	"context"
	"fmt"
	"github.com/arienmalec/alexa-go"
	ala "github.com/temesxgn/se6367-backend/alexa"
	"github.com/temesxgn/se6367-backend/auth"
	"github.com/temesxgn/se6367-backend/config"
	"github.com/temesxgn/se6367-backend/hasura"
)

// GetMyEventsForTodayIntent -
func GetMyEventsForTodayIntent(user *auth.User) (alexa.Response, error) {
	// var events []models.Event
	// user := request.Body.Intent.Slots["user"].Value

	// feedResponse, _ := RequestFeed("frontpage")
	// var builder alexa.SSMLBuilder
	// builder.Say("Here are the current frontpage deals:")
	// builder.Pause("1000")
	// for _, item := range feedResponse.Channel.Item {
	// 	builder.Say(item.Title)
	// 	builder.Pause("1000")
	// }
	// return alexa.NewSSMLResponse("Frontpage Deals", builder.Build())
	// return alexa.Response{}, nil

	var builder ala.SSMLBuilder
	service := hasura.NewService(config.GetHasuraEndpoint())
	events, _ := service.GetEvents(context.Background(), nil)
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
