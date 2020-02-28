package alexa

import (
	"strings"

	"github.com/arienmalec/alexa-go"
)

type SSML struct {
	text  string
	pause string
}

type SSMLBuilder struct {
	SSML []SSML
}

func NewSSMLResponse(title string, text string) alexa.Response {
	r := alexa.Response{
		Version: "1.0",
		Body: alexa.ResBody{
			OutputSpeech: &alexa.Payload{
				Type: "SSML",
				SSML: text,
			},
			ShouldEndSession: true,
		},
	}
	return r
}

func ParseString(text string) string {
	text = strings.ToLower(text)
	//text = strings.Replace(text, "&", "and", -1)
	//text = strings.Replace(text, "+", "plus", -1)
	//text = strings.Replace(text, "@", "at", -1)
	//text = strings.Replace(text, "w/", "with", -1)
	//text = strings.Replace(text, "in.", "inches", -1)
	//text = strings.Replace(text, "s/h", "shipping and handling", -1)
	//text = strings.Replace(text, " ac ", " after coupon ", -1)
	//text = strings.Replace(text, "fs", "free shipping", -1)
	//text = strings.Replace(text, "f/s", "free shipping", -1)
	text = strings.Replace(text, "-", "", -1)
	text = strings.Replace(text, "™", "", -1)
	text = strings.Replace(text, "  ", " ", -1)
	return text
}

func (builder *SSMLBuilder) Say(text string) {
	text = ParseString(text)
	builder.SSML = append(builder.SSML, SSML{text: text})
}

func (builder *SSMLBuilder) Pause(pause string) {
	builder.SSML = append(builder.SSML, SSML{pause: pause})
}

func (builder *SSMLBuilder) Build() string {
	var response string
	for index, ssml := range builder.SSML {
		if ssml.text != "" {
			response += ssml.text + " "
		} else if ssml.pause != "" && index != len(builder.SSML)-1 {
			response += "<break time='" + ssml.pause + "ms'/> "
		}
	}

	builder.SSML = make([]SSML, 0)
	return "<speak>" + response + "</speak>"
}
