package hasura

import (
	"context"
	"fmt"
	"github.com/temesxgn/se6367-backend/auth"
	"github.com/temesxgn/se6367-backend/config"
	"github.com/temesxgn/se6367-backend/graphql"
	"github.com/temesxgn/se6367-backend/hasura/models"
	"sync"
	"time"
)

var (
	service *hasuraService
	once    sync.Once
)

func log(s string) {
	fmt.Println("LOGGER: " + s)
}

func initialize(endpoint string) {
	once.Do(func() {
		client := graphql.NewClient(endpoint, log)
		client.AddDefaultHeader(auth.AdminSecretCtxKey.String(), config.GetHasuraSecret())
		service = &hasuraService{
			client,
		}
	})
}

func NewService(endpoint string) *hasuraService {
	initialize(endpoint)
	return service
}

type hasuraService struct {
	client *graphql.Client
}

// GetEvents - retrieve list of events based on the given filter params
func (h *hasuraService) GetEvents(ctx context.Context, filter *models.EventFilterParams) ([]*models.Event, error) {
	var respData models.GetEventsResponse
	req := graphql.NewRequest(`
		query MyEventsToday($id: String!) {
		  events(where: {account_id: { _eq: $id }}) {
			id
			title
		  }
		}
	`)

	req.Var("id", filter.UserID)
	//req.Var("date", time.Now().Format("01-02-2006"))
	err := h.client.Run(ctx, req, &respData)
	if err != nil {
		fmt.Println("ERROR " + err.Error())
		return nil, err
	}

	return respData.Data.Events, nil
}

// GetEvent - retrieve event with the given id
func (h *hasuraService) GetEvent(ctx context.Context, id string) (models.Event, error) {
	//var respData models.GetEventResponse
	//req := graphql.NewRequest(`
	//	query MyQuery($id: String!) {
	//	  event_by_pk(id: $id) {
	//		title
	//
	//	  }
	//	}
	//`)
	//
	//req.Var("id", id)
	//
	//err := h.client.Run(ctx, req, &respData)
	//if err != nil {
	//	return nil, common.NewAPIError(err.Error())
	//}
	//
	//if respData.Data == nil {
	//	return nil, common.NewNotFoundError("Event not found")
	//}
	//
	//return respData.Data, nil

	panic("event not implemented")
}

func (h *hasuraService) CreateEvent(ctx context.Context, title string, time time.Time) error {
	panic("event not implemented")
}

func (h *hasuraService) DeleteEvent(ctx context.Context, title string, day time.Time) error {
	panic("event not implemented")
}
