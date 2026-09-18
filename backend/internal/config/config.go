package config

import (
	"os"
	"watchtogether/internal/domain/errors"
)

type Config struct {
	DatabaseUrl string
	HTTPPort    string
}

func Load() (*Config, error) {
	dbUrl, err := getEnv("DATABASE_URL")
	if err != nil {
		return nil, err
	}

	port, err := getEnv("HTTP_PORT")
	if err != nil {
		return nil, err
	}

	return &Config{
		DatabaseUrl: dbUrl,
		HTTPPort:    port,
	}, nil
}

func getEnv(key string) (string, error) {
	v, ok := os.LookupEnv(key)
	if !ok || v == "" {
		return "", errors.ErrorEnvNotSet
	}
	return v, nil
}
