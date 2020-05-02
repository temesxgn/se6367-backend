package handlers

import (
	"context"
	"fmt"
	"github.com/arienmalec/alexa-go"
	ala "github.com/temesxgn/se6367-backend/alexa"
	"github.com/temesxgn/se6367-backend/auth/model"
	filters "github.com/temesxgn/se6367-backend/common/models"
	"github.com/temesxgn/se6367-backend/event"
	"gopkg.in/auth0.v3"
	"time"
)

// GetMyEventsForTodayIntent -
func GetMyEventsForTodayIntentHandler(user *model.User) (alexa.Response, error) {
	var builder ala.SSMLBuilder
	service, _ := event.GetEventService(event.HasuraEventServiceType)
	runTime := time.Now()
	fmt.Println("Getting my events for time: " + time.Now().In(time.UTC).Format(time.RFC3339))
	y, m, d := time.Now().In(time.UTC).Date()
	start := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 0, 1).Add(5 * time.Hour)
	filter := &filters.EventFilterParams{
		UserID: auth0.String(user.UserEmail()),
		From:   &start,
		To:     &end,
	}

	events, _ := service.GetEvents(context.Background(), filter)
	if len(events) == 0 {
		builder.Say("You have no events for today.")
		builder.Pause("300")
		builder.Say("To create an event, say Alexa, fair banks create event")
		return ala.NewSSMLResponse("My Events Today", builder.Build()), nil
	}

	builder.Say("Here are your events for today")
	builder.Pause("500")
	for _, e := range events {
		isPassed := e.End.Before(runTime)
		if e.IsAllDay {
			builder.Say(fmt.Sprintf("All day event"))
			builder.Pause("500")
			builder.Say(fmt.Sprintf("%s", e.Title))
		} else if !isPassed {
			builder.Say(fmt.Sprintf("%s from %s to %s", e.Title, e.Start.Add(-time.Hour*5).Format(time.Kitchen), e.End.Add(-time.Hour*5).Format(time.Kitchen)))
		}

		builder.Pause("1000")
	}

	return ala.NewSSMLResponse("My Events Today", builder.Build()), nil
}
