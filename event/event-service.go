package event

import (
	"context"
	"github.com/temesxgn/se6367-backend/common/models"
	"github.com/temesxgn/se6367-backend/config"
	"github.com/temesxgn/se6367-backend/hasura"
)

// Service - Event Service interface
type Service interface {
	GetEvents(ctx context.Context, filters *models.EventFilterParams) ([]*models.Event, error)
	GetEvent(ctx context.Context, id string) (models.Event, error)
	CreateEvent(ctx context.Context, event *models.Event) error
	DeleteEvent(ctx context.Context, id string) error
}

func GetEventService(eType ServiceType) (Service, error) {
	switch eType {
	case DBEventServiceType:
		fallthrough
	case HasuraEventServiceType:
		fallthrough
	default:
		return hasura.NewService(config.GetHasuraEndpoint()), nil
	}
}
