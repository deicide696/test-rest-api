package config

import "time"

type config struct {
	AppPort string `env:"APP_PORT" envDefault:"8080"`
	DatabaseConnection
}

type DatabaseConnection struct {
	DbHost     string        `env:"DB_HOST" envDefault:""`
	DbName     string        `env:"DB_NAME" envDefault:""`
	DbUser     string        `env:"DB_USER" envDefault:""`
	DbPassword string        `env:"DB_PASSWORD" envDefault:""`
	DbPort     string        `env:"DB_PORT" envDefault:"5432"`
	Timeout    time.Duration `env:"DB_TIMEOUT" envDefault:"10s"`
}
