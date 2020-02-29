package model

type HealthInfo struct {
	Auth0Connection    bool `json:"Auth0Connection"`
	DatabaseConnection bool `json:"DatabaseConnection"`
}
