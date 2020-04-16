package google

import (
	"errors"
	"fmt"
	"github.com/temesxgn/se6367-backend/common/models"
	"github.com/temesxgn/se6367-backend/integration/integrationtype"
	"google.golang.org/api/calendar/v3"
	"time"
)

// MapToInternalCalendar - maps google calendar model to our internal model
func MapToInternalCalendar(cal *calendar.CalendarListEntry) *models.Calendar {
	return &models.Calendar{
		Id:          cal.Id,
		Name:        cal.Summary,
		Color:       cal.BackgroundColor,
		Integration: integrationtype.GoogleServiceType,
	}
}

func MapToInternalCalendars(gCalList *calendar.CalendarList) []*models.Calendar {
	cals := make([]*models.Calendar, 0)
	for _, gCal := range gCalList.Items {
		cal := MapToInternalCalendar(gCal)
		cals = append(cals, cal)
	}

	return cals
}

func MapToInternalEvents(calID string, eventsList *calendar.Events) []*models.Event {
	events := make([]*models.Event, 0)
	for _, e := range eventsList.Items {
		event, err := MapToInternalEvent(calID, e)
		if err != nil {
			fmt.Println(fmt.Sprintf("Skipping event %s, Cause: %s", e.Id, err.Error()))
			continue
		}

		events = append(events, event)
	}

	return events
}

func MapToInternalEvent(calID string, e *calendar.Event) (*models.Event, error) {
	start, err := time.Parse(time.RFC3339, e.Start.DateTime)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error parsing start time for %v. Cause: %v", e.Id, err.Error()))
	}

	end, err := time.Parse(time.RFC3339, e.End.DateTime)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error parsing end time for %v. Cause: %v", e.Id, err.Error()))
	}

	return &models.Event{
		ID:          e.Id,
		CalendarID:  calID,
		Title:       e.Summary,
		Start:       start,
		End:         end,
		Description: e.Description,
	}, nil

}
