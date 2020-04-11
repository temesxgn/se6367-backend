package hasura

import (
	"context"
	"fmt"
	authCtx "github.com/temesxgn/se6367-backend/auth/ctx"
	"github.com/temesxgn/se6367-backend/common/client/graphql"
	"github.com/temesxgn/se6367-backend/common/models"
	"github.com/temesxgn/se6367-backend/common/util/jsonutils"
	"github.com/temesxgn/se6367-backend/config"
	"sync"
	"time"
)

var (
	service *hasuraService
	once    sync.Once
)

type hasuraService struct {
	client *graphql.Client
}

func log(s string) {
	fmt.Println("LOGGER: " + s)
}

func initialize(endpoint string) {
	once.Do(func() {
		client := graphql.NewClient(endpoint, log)
		client.AddDefaultHeader(authCtx.AdminSecretCtxKey.String(), config.GetHasuraSecret())
		service = &hasuraService{
			client,
		}
	})
}

func NewService(endpoint string) *hasuraService {
	initialize(endpoint)
	return service
}

// GetEvents - retrieve list of events based on the given filter params
func (h *hasuraService) GetEvents(ctx context.Context, filter *models.EventFilterParams) ([]*models.Event, error) {
	d, _ := jsonutils.Marshal(filter)
	fmt.Println(fmt.Sprintf("Getting events for filter: %v", d))
	var respData models.GetEventsResponse
	req := graphql.NewRequest(`
		query MyEventsToday($id: String!, $start: timestamptz, $end: timestamptz) {
		  events(where: { account_id: { _eq: $id }, start: { _gt: $start }, end: { _lte: $end }}) {
			id
			title
			start
			end
		  }
		}
	`)

	req.Var("id", filter.UserID)
	req.Var("start", filter.From.Format(time.RFC3339))
	req.Var("end", filter.To.Format(time.RFC3339))
	err := h.client.Run(ctx, req, &respData)
	if err != nil {
		fmt.Println("ERROR " + err.Error())
		return nil, err
	}

	return respData.Data.Events, nil
}

// GetEvent - retrieve event with the given id
func (h *hasuraService) GetEvent(ctx context.Context, id string) (models.Event, error) {
	//var respData model.GetEventResponse
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

func (h *hasuraService) CreateEvent(ctx context.Context, event *models.Event) error {
	user := authCtx.GetUser(ctx)
	var respData models.GetEventsResponse
	req := graphql.NewRequest(`
		mutation CreateEvent($title: String!, $account: String!, $description: String, $start: timestamptz!, $end: timestamptz!) {
			insert_event(objects: { account_id: $account, title: $title, description: $description, from: $end, start: $start,  }) {
			  returning {
				id
			  }
			}
	  	}
	`)

	req.Var("title", event.Title)
	req.Var("start", event.Start)
	req.Var("end", event.End)
	req.Var("description", event.Description)
	req.Var("account", user.Claims.XHasuraUserEmail)
	err := h.client.Run(ctx, req, &respData)
	if err != nil {
		fmt.Println("ERROR creating event " + event.Title + " " + err.Error())
		return err
	}

	return nil
}

func (h *hasuraService) DeleteEvent(ctx context.Context, id string) error {
	panic("event not implemented")
}
