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

func (h *hasuraService) GetEvents(ctx context.Context, filter models.EventFilterParams) ([]*models.Event, error) {
	var respData models.GetEventsResponse
	req := graphql.NewRequest(`
		query GetOrders($limit: Int = 10, $user: String, $customer: String, $status: order_status_enum) {
			orders(limit: $limit, where: { user_id: { _eq: $user }, stripe_customer_id: { _eq: $customer }, status: { _eq: $status }}) {
				id
				payment_intent_id
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

	req.Var("limit", filter.Limit)
	req.Var("user", filter.UserID)

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
