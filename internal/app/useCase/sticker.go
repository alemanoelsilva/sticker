package useCase

import "sticker/internal/app/entity"

func (ser *Service) CreateSticker(input entity.Sticker) error {
	ser.Logger.Info().Msg("Creating a Sticker")

	return ser.Repository.AddSticker(input)
}
