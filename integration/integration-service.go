package integration

import (
	"github.com/temesxgn/se6367-backend/common/models"
	"github.com/temesxgn/se6367-backend/integration/google"
	"github.com/temesxgn/se6367-backend/integration/integrationtype"
)

type Service interface {
	GetCalendars() ([]*models.Calendar, error)
	GetCalendarEvents(calID string) ([]*models.Event, error)
	//AddEventToCalendar(event model.Event, calendarID string) error
	//DeleteEventOnCalendar(eventID, calendarID string)
}

func GetCalendarIntegrationService(accessToken, refreshToken string, iType integrationtype.ServiceType) (Service, error) {
	switch iType {
	case integrationtype.GoogleServiceType:
		fallthrough
	default:
		return google.NewService(accessToken, refreshToken)
	}
}