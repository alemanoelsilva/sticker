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

func UpdateSticker(stickerId int, sticker *entities.Sticker) (entities.Sticker, error) {
	var stickerOnDB entities.Sticker
	database.Instance.First(&stickerOnDB, stickerId)

	if stickerOnDB.ID == 0 {
		return stickerOnDB, errors.New("Sticker not found")
	}

	stickerToUpdate := sticker
	stickerToUpdate.ID = uint(stickerId)

	database.Instance.Save(&stickerToUpdate)

	return entities.Sticker{
		ID:             stickerToUpdate.ID,
		Name:           stickerToUpdate.Name,
		Description:    stickerToUpdate.Description,
		Status:         stickerToUpdate.Status,
		IsPublic:       stickerToUpdate.IsPublic,
		IsAutoApproved: stickerToUpdate.IsAutoApproved,
	}, nil
}

func DeleteStickerById(stickerId int) (entities.Sticker, error) {
	var sticker entities.Sticker
	database.Instance.First(&sticker, stickerId)

	if sticker.ID == 0 {
		return sticker, errors.New("Sticker not found")
	}

	database.Instance.Delete(&sticker, stickerId)
	return sticker, nil
}
