package handlers

import (
	"github.com/arienmalec/alexa-go"
	ala "github.com/temesxgn/se6367-backend/alexa"
)

func HandleHelpIntent(request alexa.Request) (alexa.Response, error) {
	var builder ala.SSMLBuilder
	builder.Say("Here are some of the things you can ask:")
	builder.Pause("1000")
	builder.Say("What are my events for today.")
	builder.Pause("1000")
	builder.Say("Give me the popular deals.")
	return ala.NewSSMLResponse("Help", builder.Build()), nil
}
