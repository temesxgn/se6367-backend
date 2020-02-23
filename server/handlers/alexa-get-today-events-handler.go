package handlers

import (
	"context"

	"github.com/arienmalec/alexa-go"
)

// GetMyEventsForTodayIntent -
func GetMyEventsForTodayIntent(ctx context.Context, request alexa.Request) (alexa.Response, error) {
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
	return alexa.Response{}, nil
}
