package useCase

import (
	Repository "sticker/internal/app/repository/mysql"

	"github.com/rs/zerolog"
)

type Service struct {
	Repository Repository.MysqlDB
	Logger     *zerolog.Logger
}

func LoadService(repository Repository.MysqlDB, logger *zerolog.Logger) *Service {
	return &Service{
		Logger:     logger,
		Repository: repository,
	}
}
