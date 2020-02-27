package event

import (
	"context"
	"github.com/temesxgn/se6367-backend/config"
	"github.com/temesxgn/se6367-backend/hasura"
	"time"

	"github.com/temesxgn/se6367-backend/hasura/models"
)

// Service - Event Service interface
type Service interface {
	GetEvents(ctx context.Context, filters *models.EventFilterParams) ([]*models.Event, error)
	GetEvent(ctx context.Context, title string) (models.Event, error)
	CreateEvent(ctx context.Context, title string, time time.Time) error
	DeleteEvent(ctx context.Context, title string, day time.Time) error
}

func GetEventService(eType ServiceType) Service {
	switch eType {
	case DBServiceType:
		fallthrough
	case HasuraServiceType:
		fallthrough
	default:
		return hasura.NewService(config.GetHasuraEndpoint())
	}
}
