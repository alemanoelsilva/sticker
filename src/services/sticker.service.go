package services

import (
	"errors"
	"sticker/src/config/database"
	"sticker/src/entities"
)

func GetStickers() []entities.Sticker {
	var stickers []entities.Sticker
	database.Instance.Find(&stickers)
	return stickers
}

func GetStickerById(stickerId string) (entities.Sticker, error) {
	var sticker entities.Sticker
	database.Instance.First(&sticker, stickerId)

	if sticker.ID == 0 {
		return sticker, errors.New("Sticker not found")
	}

	return sticker, nil
}

func CreateSticker(sticker *entities.Sticker) {
	database.Instance.Create(&sticker)
}
