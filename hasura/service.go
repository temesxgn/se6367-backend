package hasura

import (
	"context"

	"github.com/temesxgn/se6367-backend/hasura/models"
)

// Service - Hasura Service interface
type Service interface {
	GetEvents(ctx context.Context, filters *models.EventFilterParams) ([]*models.Event, error)
	GetEvent(ctx context.Context, id string) (models.Event, error)
}

// NewService - creates a new instance
func NewService(endpoint string) Service {
	return NewDefaultService(endpoint)
}
