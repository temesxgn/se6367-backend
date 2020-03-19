package google

import (
	"context"
	"fmt"
	"github.com/temesxgn/se6367-backend/common/models"
	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

func NewService(accessToken, refreshToken string) (*googleService, error) {
	fmt.Println(fmt.Sprintf("ACCESS TOKEN: %v", accessToken))
	fmt.Println(fmt.Sprintf("REFRESH TOKEN: %v", refreshToken))
	tkn := &oauth2.Token{
		AccessToken:  accessToken,
		TokenType:    "Bearer",
		RefreshToken: refreshToken,
	}

	tokenSource := oauth2.ReuseTokenSource(tkn, oauth2.StaticTokenSource(tkn))
	service, err := calendar.NewService(context.Background(), option.WithTokenSource(tokenSource))
	if err != nil {
		fmt.Println("Error creating client for google: " + err.Error())
		return nil, err
	}

	return &googleService{
		service,
	}, nil
}

type googleService struct {
	service *calendar.Service
}

// GetCalendars - retrieve the list of users calendar
func (s *googleService) GetCalendars() ([]*models.Calendar, error) {
	calList, err := s.service.CalendarList.List().Do()
	if err != nil {
		fmt.Println(fmt.Sprintf("ERROR get calendars from google %v", err.Error()))
		return make([]*models.Calendar, 0), err
	}

	cals := MapToInternalCalendars(calList)
	return cals, nil
}

func (s *googleService) GetCalendarEvents(calID string) ([]*models.Event, error) {
	eventList, err := s.service.Events.List(calID).Do()
	if err != nil {
		fmt.Println("ERROR getting calendar events for cal" + calID + " from google " + err.Error())
		return make([]*models.Event, 0), err
	}

	events := MapToInternalEvents(calID, eventList)
	return events, nil
}