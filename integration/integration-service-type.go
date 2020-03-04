package integration

type ServiceType string

const (
	GoogleServiceType ServiceType = "google"
)

func (s ServiceType) String() string {
	return string(s)
}