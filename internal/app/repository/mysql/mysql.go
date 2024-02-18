package mysql

import (
	"sticker/config"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type MysqlDB struct {
	db *sqlx.DB
}

func NewMysqlDB(db *sqlx.DB) *MysqlDB {
	return &MysqlDB{
		db: db,
	}
}

func RunMigrations(db *sqlx.DB, logger *zerolog.Logger) error {
	driver, err := mysql.WithInstance(db.DB, &mysql.Config{})
	if err != nil {
		return err
	}

	url := config.AppConfig.MigrationPath

	m, err := migrate.NewWithDatabaseInstance(url, "mysql", driver)
	if err != nil {
		logger.Fatal().Msg(err.Error())
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		logger.Fatal().Msg(err.Error())
		return err
	}

	return nil
}
