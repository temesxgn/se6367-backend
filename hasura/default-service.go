package hasura

import (
	"context"
	"sync"

	"github.com/temesxgn/se6367-backend/auth"
	"github.com/temesxgn/se6367-backend/common"
	"github.com/temesxgn/se6367-backend/config"
	"github.com/temesxgn/se6367-backend/graphql"
	"github.com/temesxgn/se6367-backend/hasura/models"
)

var (
	service *hasuraService
	once    sync.Once
)

func initialize(endpoint string) {
	once.Do(func() {
		client := graphql.NewClient(endpoint)
		client.AddDefaultHeader(auth.AdminSecretCtxKey.String(), config.GetHasuraSecret())
		service = &hasuraService{
			client,
		}
	})
}

func NewDefaultService(endpoint string) *hasuraService {
	initialize(endpoint)
	return service
}

// HasuraService - Service to dispatch actions to Hasura Service
type hasuraService struct {
	client *graphql.Client
}

func (h *hasuraService) GetEvents(ctx context.Context, filter *models.EventFilterParams) ([]*models.Event, error) {
	var respData models.GetEventsResponse
	req := graphql.NewRequest(`
		query MyQuery {
		  event(where: {account_id: {_eq: "auth0|5dd98f908537f90eefda947d"}}, limit: 3) {
			id
			title
			description
		  }
		}
	`)

	err := h.client.Run(ctx, req, &respData)
	if err != nil {
		return nil, err
	}

	return respData.Data, nil
}

// GetEvent - retrieve event with the given id
func (h *hasuraService) GetEvent(ctx context.Context, id string) (*models.Event, error) {
	var respData models.GetEventResponse
	req := graphql.NewRequest(`
		query GetOrder($id: String!) {
			orders_by_pk(id: $id) {
				id
				payment_intent_amount
				refund_percent
				status
				created_at
				order_items {
					amount
					product_id
					quantity
				}
			}
		}
	`)

	req.Var("id", id)

	err := h.client.Run(ctx, req, &respData)
	if err != nil {
		return nil, common.NewAPIError(err.Error())
	}

	if respData.Data == nil {
		return nil, common.NewNotFoundError("Event not found")
	}

	return respData.Data, nil
}
