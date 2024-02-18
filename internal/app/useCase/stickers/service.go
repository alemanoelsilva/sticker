package stickers

import (
	repo "sticker/internal/app/repository/mysql/stickers"

	"github.com/rs/zerolog"
)

type Service struct {
	Repository repo.SqlRepository
	Logger     *zerolog.Logger
}

func LoadService(repository repo.SqlRepository, logger *zerolog.Logger) *Service {
	return &Service{
		Logger:     logger,
		Repository: repository,
	}
}
