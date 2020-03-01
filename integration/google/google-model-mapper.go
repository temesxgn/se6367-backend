package google

import (
	"github.com/temesxgn/se6367-backend/common/models"
	"google.golang.org/api/calendar/v3"
)


// MapToInternalCalendar - maps google calendar model to our internal model
func MapToInternalCalendar(cal *calendar.Calendar) *models.Calendar {
	return &models.Calendar{Id: cal.Id}
}

func MapToInternalCalendarList(gCalList *calendar.CalendarList) []*models.Calendar {
	cals := make([]*models.Calendar, 0)

	for _, gCal := range gCalList.Items {
		cal := &models.Calendar{Id: gCal.Id}
		cals = append(cals, cal)
	}

	return cals
}

func MapToInternalEvents(eventsList *calendar.Events) []*models.Event {
	events := make([]*models.Event, 0)

	for _, e := range eventsList.Items {
		event := &models.Event{
			ID:          e.Id,
			Title:      e.Summary,
			Description: e.Description,
		}

		events = append(events, event)
	}

	return events
}

func MapToInternalEvent(event *calendar.Event) *models.Event {
	return &models.Event{
		ID:          event.Id,
		Title:       event.Summary,
		Description: event.Description,
	}
}