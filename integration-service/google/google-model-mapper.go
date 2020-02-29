package google

import (
	"github.com/temesxgn/se6367-backend/common/models"
	calendar "google.golang.org/api/calendar/v3"
)


// MapToInternalCalendar - maps google calendar model to our internal model
func MapToInternalCalendar(cal calendar.Calendar) *models.Calendar {

}

func MapToInternalCalendarList(cal *calendar.CalendarList) []*models.Calendar {

}
