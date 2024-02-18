package database

import (
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

type DDB struct {
	Logger *zerolog.Logger
}

func (ddb *DDB) Connect(uri string) *sqlx.DB {
	db, err := sqlx.Connect("mysql", uri)
	if err != nil {
		ddb.Logger.Fatal().Err(err).Msg("MySQL connection error")
		os.Exit(1)
	}

	return db
}

func (ddb *DDB) RunMigrations() {
	ddb.Logger.Info().Msg("Not Implemented yet")
}
