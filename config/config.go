package config

import "os"

func GetEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// / GO_ENV: Should be one of "dev", "prod"
var GO_ENV = GetEnvWithDefault("GO_ENV", "dev")

var ProductionMode = GO_ENV == "prod"
