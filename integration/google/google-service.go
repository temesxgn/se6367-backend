package google

import (
	"context"
	"fmt"
	"github.com/temesxgn/se6367-backend/common/models"
	calendar "google.golang.org/api/calendar/v3"
)

func NewService(token string) (*googleService, error) {
	service, err := calendar.NewService(context.Background())
	if err != nil {
		fmt.Println("Error creating client for google: " + err.Error())
		return nil, err
	}

	return &googleService{
		service,
		token,
	}, nil
}

type googleService struct {
	service *calendar.Service
	token string
}

// GetCalendars - retrieve the list of users calendar
func (s *googleService) GetCalendars() ([]*models.Calendar, error) {
	calList, err := s.service.CalendarList.List().Do()
	if err != nil {
		fmt.Println("ERROR get calendars from google " + err.Error())
		return nil, err
	}

	cals := MapToInternalCalendarList(calList)
	return cals, nil
}

func (s *googleService) GetCalendarEvents(calID string) ([]*models.Event, error) {
	eventList, err := s.service.Events.List(calID).Do()
	if err != nil {
		fmt.Println("ERROR getting calendar events for cal" + calID + " from google " + err.Error())
		return nil, err
	}

	events := MapToInternalEvents(eventList)
	return events, nil
}