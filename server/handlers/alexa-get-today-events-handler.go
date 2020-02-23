package handlers

import (
	"github.com/arienmalec/alexa-go"
	ala "github.com/temesxgn/se6367-backend/alexa"
)

// GetMyEventsForTodayIntent -
func GetMyEventsForTodayIntent(request *alexa.Request) (alexa.Response, error) {
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
	builder.Say("Here are your events for today")
	builder.Pause("500")
	builder.Say("Doctor's appointment at Nine AM until four PM")

	builder.Pause("1000")
	builder.Say("Gym at eleven thirty AM until one pm")

	return ala.NewSSMLResponse("My Events Today", builder.Build()), nil

}
