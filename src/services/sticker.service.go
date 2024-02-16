package services

import (
	"errors"
	"sticker/src/config/database"
	stickers_entity "sticker/src/entities/stickers"
)

func GetStickers() []stickers_entity.Sticker {
	var stickers []stickers_entity.Sticker
	database.Instance.Find(&stickers)
	return stickers
}

func GetStickerById(stickerId string) (stickers_entity.Sticker, error) {
	var sticker stickers_entity.Sticker
	database.Instance.First(&sticker, stickerId)

	if sticker.ID == 0 {
		return sticker, errors.New("Sticker not found")
	}

	return sticker, nil
}

func CreateSticker(sticker *stickers_entity.Sticker) {
	database.Instance.Create(&sticker)
}

func UpdateSticker(stickerId int, sticker *stickers_entity.Sticker) (stickers_entity.Sticker, error) {
	var stickerOnDB stickers_entity.Sticker
	database.Instance.First(&stickerOnDB, stickerId)

	if stickerOnDB.ID == 0 {
		return stickerOnDB, errors.New("Sticker not found")
	}

	stickerToUpdate := sticker
	stickerToUpdate.ID = uint(stickerId)

	database.Instance.Save(&stickerToUpdate)

	return stickers_entity.Sticker{
		ID:             stickerToUpdate.ID,
		Name:           stickerToUpdate.Name,
		Description:    stickerToUpdate.Description,
		Status:         stickerToUpdate.Status,
		IsPublic:       stickerToUpdate.IsPublic,
		IsAutoApproved: stickerToUpdate.IsAutoApproved,
	}, nil
}

func DeleteStickerById(stickerId int) (stickers_entity.Sticker, error) {
	var sticker stickers_entity.Sticker
	database.Instance.First(&sticker, stickerId)

	if sticker.ID == 0 {
		return sticker, errors.New("Sticker not found")
	}

	database.Instance.Delete(&sticker, stickerId)
	return sticker, nil
}
