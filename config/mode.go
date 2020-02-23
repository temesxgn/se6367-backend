package config

// AppMode - app mode
type AppMode string

// possible app modes
const (
	LocalMode AppMode = "local"
	DevMode   AppMode = "develop"
	QAMode    AppMode = "qa"
	ProdMode  AppMode = "production"
)

func (m AppMode) String() string {
	return string(m)
}

// IsProductionMode - check if mode is set as production
func (m AppMode) IsProductionMode() bool {
	return m == ProdMode
}
