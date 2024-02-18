package config

import (
	"os"

	"github.com/rs/zerolog"
)

type EnvConfig struct {
	Port             string
	ConnectionString string
	SecretToken      string
	MigrationPath    string
}

var AppConfig *EnvConfig

func LoadAppConfig(logger *zerolog.Logger) {
	logger.Info().Msg("Loading Server Configurations...")

	envPort := os.Getenv("PORT")
	envDBString := os.Getenv("DB_STRING")
	envJwtToken := os.Getenv("JWT_TOKEN")
	envMigrationPath := os.Getenv("MIGRATION_PATH")

	AppConfig = &EnvConfig{
		Port:             envPort,
		ConnectionString: envDBString,
		SecretToken:      envJwtToken,
		MigrationPath:    envMigrationPath,
	}

}
