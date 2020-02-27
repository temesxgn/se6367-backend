package event

type ServiceType string

const (
	HasuraServiceType ServiceType = "hasura"
	DBServiceType ServiceType = "database"
)