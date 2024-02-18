package stickers

import (
	"sticker/internal/app/entity"
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
func (ser *Service) CreateSticker(input entity.Sticker, userId int) error {
	ser.Logger.Info().Msg("Creating a Sticker")

	return ser.Repository.AddSticker(input, userId)
}

func (ser *Service) GetStickers(userId int) ([]entity.Sticker, error) {
	ser.Logger.Info().Msg("Getting a Stickers")

	return ser.Repository.GetStickers(userId)
}

func (ser *Service) GetStickerById(userId int, stickerId int) (entity.Sticker, error) {
	ser.Logger.Info().Msg("Getting Sticker by id")

	return ser.Repository.GetStickerById(userId, stickerId)
}
