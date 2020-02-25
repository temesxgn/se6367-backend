package hasura

import (
	"context"
	"fmt"
	"github.com/temesxgn/se6367-backend/util/jsonutils"
	"gopkg.in/auth0.v3"
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
		query MyQuery($id: String!) {
		  event(where: {account_id: {_eq: $id}}, limit: 3) {
			id
			title
			description
		  }
		}
	`)

	fmt.Println("SEACHING USER " + auth0.StringValue(filter.UserID))
	req.Var("id", filter.UserID)
	err := h.client.Run(ctx, req, &respData)
	if err != nil {
		fmt.Println("ERROR " + err.Error())
		return nil, err
	}

	dt, _ := jsonutils.Marshal(respData)
	fmt.Println(fmt.Sprintf("RES: %v", dt))
	return respData.Data, nil
}

// GetEvent - retrieve event with the given id
func (h *hasuraService) GetEvent(ctx context.Context, id string) (*models.Event, error) {
	var respData models.GetEventResponse
	req := graphql.NewRequest(`
		query MyQuery($id: String!) {
		  event_by_pk(id: $id) {
			title
			
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
