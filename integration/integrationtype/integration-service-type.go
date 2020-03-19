package integrationtype

type ServiceType string

const (
	GoogleServiceType ServiceType = "GOOGLE"
)

func (s ServiceType) String() string {
	return string(s)
}