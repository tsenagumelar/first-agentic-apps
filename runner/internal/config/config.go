package config

import (
	"os"
	"strconv"
	"time"
)

type Config struct {
	AppPort     string
	CodexBinary string
	Timeout     time.Duration
}

func Load() Config {
	timeoutSeconds := getEnvAsInt("CODEX_TIMEOUT_SECONDS", 300)

	return Config{
		AppPort:     getEnv("APP_PORT", "7001"),
		CodexBinary: getEnv("CODEX_BINARY", "codex"),
		Timeout:     time.Duration(timeoutSeconds) * time.Second,
	}
}

func getEnv(key string, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}

	return parsed
}
