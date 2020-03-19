// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package graph

import (
	"context"
	"errors"
	"fmt"
	"github.com/temesxgn/se6367-backend/event"

	"github.com/temesxgn/se6367-backend/auth"
	authCtx "github.com/temesxgn/se6367-backend/auth/ctx"
	integrationService "github.com/temesxgn/se6367-backend/integration"
	"github.com/temesxgn/se6367-backend/integration/integrationtype"
)

func (r *mutationResolver) SyncEvents(ctx context.Context, integration integrationtype.ServiceType) (bool, error) {
	user := authCtx.GetUser(ctx)
	authService, err := auth.GetAuthService(auth.AuthZeroAuthServiceType)
	if err != nil {
		return false, errors.New(fmt.Sprintf("Error getting auth service. Cause: %v", err.Error()))
	}

	usr, err := authService.GetUser(user.UserID())
	if err != nil {
		return false, errors.New(fmt.Sprintf("Error getting auth user. Cause: %v", err.Error()))
	}

	accessToken, refreshToken := usr.GetIdentityProviderTokens(integration)
	svc, err := integrationService.GetCalendarIntegrationService(accessToken, refreshToken, integration)
	if err != nil {
		return false, errors.New(fmt.Sprintf("Error getting calendar integration service. Cause: %v", err.Error()))
	}

	cals, err := svc.GetCalendars()
	if err != nil {
		return false, errors.New(fmt.Sprintf("Error getting calendars. Cause: %v", err.Error()))
	}

	eventService, err := event.GetEventService(event.HasuraEventServiceType)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error getting event service. Cause: %v", err.Error()))
		return false, err
	}

	for _, cal := range cals {
		events, err := svc.GetCalendarEvents(cal.Id)
		if err != nil {
			fmt.Println(fmt.Sprintf("Error getting calendar %v events. Cause: %v", cal.Id, err.Error()))
			continue
		}

		for _, e := range events {
			if err := eventService.CreateEvent(ctx, e); err != nil {
				fmt.Println(fmt.Sprintf("Failed creating event %v", err.Error()))
			}
		}
	}

	return true, nil
}
