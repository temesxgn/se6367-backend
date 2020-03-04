// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
package graph

import (
	"context"
	"fmt"
	"github.com/temesxgn/se6367-backend/auth"

	ctx2 "github.com/temesxgn/se6367-backend/auth/ctx"
	integration2 "github.com/temesxgn/se6367-backend/integration"
)

func (r *mutationResolver) SyncEvents(ctx context.Context, integration integration2.ServiceType) (bool, error) {
	user := ctx2.GetUser(ctx)
	authService, err := auth.GetAuthService(auth.AuthZeroAuthServiceType)
	if err != nil {
		fmt.Println(err.Error())
		return false, nil
	}

	usr, err := authService.GetUser(user.UserID())
	if err != nil {
		fmt.Println(err.Error())
		return false, nil
	}

	svc, err := integration2.GetCalendarIntegrationService(usr.GetIdentityProviderAccessToken(integration.String()), integration)
	if err != nil {
		fmt.Println(err.Error())
		return false, nil
	}

	cals, err := svc.GetCalendars()
	if err != nil {
		fmt.Println(err.Error())
		return false, nil
	}

	for _, cal := range cals {
		fmt.Println(cal.Id)
	}

	return true, nil
}
