package config

import (
	"os"
	"strconv"
)

type Config struct {
	HTTPAddr          string
	GRPCAddr          string
	DBPath            string
	CORSOrigin        string
	MaxMessages       int
	MaxRequestBody    int
	MaxCiphertextSize int
	MaxRecipients     int
}

func Load() Config {
	return Config{
		HTTPAddr:          envOrDefault("NCHAT_HTTP_ADDR", ":8080"),
		GRPCAddr:          envOrDefault("NCHAT_GRPC_ADDR", ":9090"),
		DBPath:            envOrDefault("NCHAT_DB_PATH", "./data/nchat.db"),
		CORSOrigin:        envOrDefault("NCHAT_CORS_ORIGIN", "*"),
		MaxMessages:       envIntOrDefault("NCHAT_MAX_MESSAGES", 200),
		MaxRequestBody:    envIntOrDefault("NCHAT_MAX_REQUEST_BODY_BYTES", 10*1024*1024),
		MaxCiphertextSize: envIntOrDefault("NCHAT_MAX_CIPHERTEXT_CHARS", 14*1024*1024),
		MaxRecipients:     envIntOrDefault("NCHAT_MAX_RECIPIENTS", 64),
	}
}

func envOrDefault(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}
	return value
}

func envIntOrDefault(key string, fallback int) int {
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
