package config

import (
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv string `env:"APP_ENV" env-default:"local"`
	DbConf DBConfig
}

type DBConfig struct {
	Host            string        `env:"DB_HOST" env-required:"true"`
	Port            int           `env:"DB_PORT" env-default:"5432"`
	User            string        `env:"DB_USER" env-required:"true"`
	Password        string        `env:"DB_PASSWORD" env-required:"true"`
	DBName          string        `env:"DB_NAME" env-required:"true"`
	SSLMode         string        `env:"DB_SSLMODE" env-default:"disable"`
	MaxConns        int32         `env:"DB_MAX_CONNS" env-default:"10"`
	MinConns        int32         `env:"DB_MIN_CONNS" env-default:"2"`
	MaxConnLifetime time.Duration `env:"DB_MAX_CONN_LIFETIME" env-default:"30m"`
	MaxConnIdleTime time.Duration `env:"DB_MAX_CONN_IDLE_TIME" env-default:"5m"`
}

func Load() (*Config, error) {
	_ = godotenv.Load()

	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, ErrCantReadConfig
	}
	return &cfg, nil
}
