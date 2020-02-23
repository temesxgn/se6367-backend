package metrics

// TODO move out to seperate repo

// Actions
const (
	BackendServer = "backend.server"
)

// Tags
const (
	CreateEvent = "create.event"
)

// Logger WithField Stats Keys
const (
	Mode            = "mode"
	AppName         = "app.name"
	TransactionType = "txn.type"
	Request         = "request"
	Response        = "response"
	EventID         = "event.id"
)
