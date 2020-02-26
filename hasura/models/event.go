package models

import "time"

// Event - event model object
type Event struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type EventFilterParams struct {
	UserID *string    `json:"user_id"`
	To     *time.Time `json:"to"`
	From   *time.Time `json:"from"`
	Limit  *int       `json:"limit"`
}

// GetEventsResponse - list of event response model
type GetEventsResponse struct {
	Data EventsResponseData `json:"data"`
}

type EventsResponseData struct {
	Events []*Event `json:"events"`
}

// GetEventResponse - event response model
type GetEventResponse struct {
	Data Event `json:"event_by_pk"`
}
