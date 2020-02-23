package config

import (
	"os"
	"strconv"
)

// DefaultPort - Default server port
const (
	DefaultPort             = "8081"
	DefaultStripeRetryCount = 3
)

// GetValue - loads the config value from the internal map, if not presents loads it from the env, if still missing uses the default
func GetValue(key string, defaultVal string) string {
	return GetWithDefault(key, defaultVal).(string)
}

// GetIntValue - loads the config value from the internal map, if not presents loads it from the env, if still missing uses the default
func GetIntValue(key string, defaultVal int64) int64 {
	val := GetWithDefault(key, defaultVal).(string)
	iv, e2 := strconv.Atoi(val)
	if e2 != nil {
		return defaultVal
	}

	return int64(iv)
}

// GetWithDefault - get key with default value if not found
func GetWithDefault(key string, defaultVal interface{}) interface{} {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}

	return val
}

// GetHasuraSecret - returns the Hasura secret or the default
func GetHasuraSecret() string {
	return GetValue(HasuraSecretKey, "")
}

// GetHasuraEndpoint - returns the Hasura endpoint
func GetHasuraEndpoint() string {
	return GetValue(HasuraEndpointKey, "")
}

// GetApplicationMode - returns the application mode or default of DEV
func GetApplicationMode() AppMode {
	return AppMode(GetValue(ApplicationModeKey, DevMode.String()))
}

// GetServerPort - port for server to run
func GetServerPort() string {
	return GetValue(ServerPort, DefaultPort)
}
