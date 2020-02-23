package alexa

type IntentType string

const (
	GetMyEventsForTodayIntentType IntentType = "GetMyEventsForTodayIntent"
	HelpIntentType                IntentType = "HelpIntent"
)

func (i IntentType) String() string {
	return string(i)
}
