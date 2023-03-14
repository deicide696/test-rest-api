package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v7"
	"github.com/sirupsen/logrus"
)

var Config config

func init() {

	if err := env.Parse(&Config); err != nil {
		logrus.Fatalf("Error initializing: %s", err.Error())
	}

	loc, err := time.LoadLocation(Config.DefaultTimezone)
	if err != nil {
		logrus.Fatalf("Invalid timezone: %s", Config.DefaultTimezone)
	}

	Config.DefaultLocation = loc
}

type config struct {
	AppPort         string         `env:"APP_PORT" envDefault:"8080"`
	DefaultTimezone string         `env:"DEFAULT_TIMEZONE" envDefault:"America/Bogota"`
	DefaultLocation *time.Location `env:"-"`
	DatabaseConnection
	Environment string `env:"ENV_APP" envDefault:"LOCAL"`
}

type DatabaseConnection struct {
	DbHost     string        `env:"DB_HOST" envDefault:""`
	DbUser     string        `env:"DB_USER" envDefault:""`
	DbPassword string        `env:"DB_PASSWORD" envDefault:""`
	DbName     string        `env:"DB_NAME" envDefault:""`
	DbPort     string        `env:"DB_PORT" envDefault:"5432"`
	Timeout    time.Duration `env:"DB_TIMEOUT" envDefault:"10s"`
}

func (dc *DatabaseConnection) DBConnectionString() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=false timezone=%s",
		dc.DbHost,
		dc.DbUser,
		dc.DbPassword,
		dc.DbName,
		dc.DbPort,
		Config.DefaultTimezone,
	)
}
