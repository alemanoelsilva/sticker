package stickers

import "sticker/internal/app/entity"

func (ser *Service) CreateSticker(input entity.Sticker, userId int) error {
	ser.Logger.Info().Msg("Creating a Sticker")

	return ser.Repository.AddSticker(input, userId)
}

func (ser *Service) GetStickers(userId int) ([]entity.Sticker, error) {
	ser.Logger.Info().Msg("Getting a Stickers")

	return ser.Repository.GetStickers(userId)
}
