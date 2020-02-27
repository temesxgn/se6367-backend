package alexa

// IntentType - alexa intent types
type IntentType string

// possible alexa intents
const (
	GetMyEventsForTodayIntentType IntentType = "GetMyEventsForTodayIntent"
	GetEventsForDayIntentType     IntentType = "GetEventsForDayIntent"
	CreateEventIntentType         IntentType = "CreateEventIntent"
	DeleteEventIntentType         IntentType = "DeleteEventIntent"
	SyncEventsIntenType           IntentType = "SyncEventsIntent"
	HelpIntentType                IntentType = "HelpIntent"
)

func (i IntentType) String() string {
	return string(i)
}
