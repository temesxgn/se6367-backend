package handlers

import (
	"context"

	"github.com/arienmalec/alexa-go"
)

// IntentDispatcher -
func IntentDispatcher(ctx context.Context, request alexa.Request) (alexa.Response, error) {
	switch request.Body.Intent.Name {
	case "GetMyEventsForTodayIntent":
		return GetMyEventsForTodayIntent(ctx, request)
	default:
		return HandleHelpIntent(request)
	}
}
