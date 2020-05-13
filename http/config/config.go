package config

import (
	"os"
)

type Config struct {
	Http_port string
	Redis_url string
}

func Load() *Config {
	return &Config{
		Http_port: getEnv("HTTP_PORT", "8686"),
		Redis_url: getEnv("REDIS_URL", "localhost:6379"),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
