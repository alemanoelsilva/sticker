package database

import (
	"os"
	"sticker/internal/app/repository/mysql/model"

	"github.com/rs/zerolog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DDB struct {
	Logger *zerolog.Logger
}

func (ddb *DDB) Connect(uri string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(uri), &gorm.Config{})

	if err != nil {
		ddb.Logger.Fatal().Err(err).Msg("MySQL connection error")
		os.Exit(1)
	}

	return db
}

func (ddb *DDB) RunMigrations(db *gorm.DB) {
	db.AutoMigrate(&model.User{}, &model.Sticker{})
}
