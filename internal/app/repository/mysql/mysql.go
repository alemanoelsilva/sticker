package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"

	sticker "sticker/internal/app/repository/mysql/stickers"
	user "sticker/internal/app/repository/mysql/users"

	_ "github.com/go-sql-driver/mysql"
)

func NewSqlRepository(db *sqlx.DB, logger *zerolog.Logger) (*user.SqlRepository, *sticker.SqlRepository) {
	userRepo := user.SqlRepository{DB: db, Logger: logger}
	stickerRepo := sticker.SqlRepository{DB: db, Logger: logger}

	return &userRepo, &stickerRepo
}
