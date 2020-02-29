package integration_service

import (
	"github.com/temesxgn/se6367-backend/common/models"
	"github.com/temesxgn/se6367-backend/integration-service/google"
)

type Service interface {
	GetCalendars() ([]*models.Calendar, error)
	GetCalendarEvents(calID string) ([]*models.Event, error)
	//AddEventToCalendar(event models.Event, calendarID string) error
	//DeleteEventOnCalendar(eventID, calendarID string)
}

func GetIntegrationService(token string, iType ServiceType) (Service, error) {
	switch iType {
	case GoogleServiceType:
		fallthrough
	default:
		return google.NewService(token)
	}
}