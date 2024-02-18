package mysql

import (
	"github.com/jmoiron/sqlx"

	sticker "sticker/internal/app/repository/mysql/stickers"
	user "sticker/internal/app/repository/mysql/users"

	_ "github.com/go-sql-driver/mysql"
)

func NewSqlRepository(db *sqlx.DB) (*user.SqlRepository, *sticker.SqlRepository) {
	userRepo := user.SqlRepository{DB: db}
	stickerRepo := sticker.SqlRepository{DB: db}

	return &userRepo, &stickerRepo
}
