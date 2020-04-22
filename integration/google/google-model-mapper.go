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
	if e.Start.DateTime != "" && e.End.DateTime != "" {
		start, startErr := time.Parse(time.RFC3339, e.Start.DateTime)
		end, endErr := time.Parse(time.RFC3339, e.End.DateTime)
		if startErr != nil && endErr != nil {
			return nil, errors.New(fmt.Sprintf("Error parsing start & end datetimes for %v. Cause: %v %v", e.Id, startErr.Error(), endErr.Error()))
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

	if e.Start.Date != "" && e.End.Date != "" {
		start, startErr := time.Parse(time.RFC3339, e.Start.Date)
		end, endErr := time.Parse(time.RFC3339, e.End.Date)
		if startErr != nil && endErr != nil {
			return nil, errors.New(fmt.Sprintf("Error parsing start & end times for %v. Cause: %v %v", e.Id, startErr.Error(), endErr.Error()))
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

	return nil, errors.New(fmt.Sprintf("Error parsing date for cal -> %v event -> %v (id: %v)"+calID, e.Summary, e.Id))
}
