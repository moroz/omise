package config

import (
	"log"
	"os"
)

func GetEnvWithDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func MustGetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s is not set!\n", key)
	}
	return value
}

// GO_ENV: Should be one of "dev", "prod"
var GO_ENV = GetEnvWithDefault("GO_ENV", "dev")
var DATABASE_URL = MustGetEnv("DATABASE_URL")

var ProductionMode = GO_ENV == "prod"
