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
	"time"
)

// GetMyEventsForTodayIntent -
func GetMyEventsForTodayIntentHandler(user *model.User) (alexa.Response, error) {
	var builder ala.SSMLBuilder
	service, _ := event.GetEventService(event.HasuraEventServiceType)
	y, m, d := time.Now().In(time.UTC).Date()
	start := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 0, 1).Add(-1 * time.Second)
	filter := &models2.EventFilterParams{
		UserID: auth0.String(user.UserEmail()),
		From:   &start,
		To:     &end,
	}
	events, _ := service.GetEvents(context.Background(), filter)

	if len(events) == 0 {
		builder.Say("You have no events for today.")
		builder.Pause("300")
		builder.Say("To create an event, say Alexa, tia mo create event")
	} else {
		builder.Say("Here are your events for today")
		builder.Pause("500")
		for _, event := range events {
			builder.Say(fmt.Sprintf("%s at 9:30 pm", event.Title))
			builder.Pause("1000")
		}
	}

	return ala.NewSSMLResponse("My Events Today", builder.Build()), nil
}
