package event

type ServiceType string

const (
	HasuraEventServiceType ServiceType = "hasura"
	DBEventServiceType     ServiceType = "database"
)