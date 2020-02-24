package alexa

// IntentType - alexa intent types
type IntentType string

// possible alexa intents
const (
	GetMyEventsForTodayIntentType IntentType = "GetMyEventsForTodayIntent"
	CreateEventIntentType         IntentType = "CreateEventIntentType"
	DeleteEventIntentType         IntentType = "DeleteEventIntentType"
	SyncEventsIntenType           IntentType = "SyncEventsIntentType"
	HelpIntentType                IntentType = "HelpIntent"
)

func (i IntentType) String() string {
	return string(i)
}
