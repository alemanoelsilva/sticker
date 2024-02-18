package config

import (
	"os"
	"strconv"

	"github.com/rs/zerolog"
)

type EnvConfig struct {
	Port               string
	ConnectionString   string
	SecretToken        string
	MigrationPath      string
	MinutesToJwtExpire int
}

var AppConfig *EnvConfig

func LoadAppConfig(logger *zerolog.Logger) {
	logger.Info().Msg("Loading Server Configurations...")

	envPort := os.Getenv("PORT")
	envDBString := os.Getenv("DB_STRING")
	envJwtToken := os.Getenv("JWT_TOKEN")
	envMigrationPath := os.Getenv("MIGRATION_PATH")

	envMinutesToJwtExpire, err := strconv.Atoi(os.Getenv("MINUTES_TO_JWT_EXPIRE"))
	if err != nil {
		logger.Fatal().Msg("Invalid or missing MINUTES_TO_JWT_EXPIRE")
		os.Exit(1)
	}

	AppConfig = &EnvConfig{
		Port:               envPort,
		ConnectionString:   envDBString,
		SecretToken:        envJwtToken,
		MigrationPath:      envMigrationPath,
		MinutesToJwtExpire: envMinutesToJwtExpire,
	}

}
