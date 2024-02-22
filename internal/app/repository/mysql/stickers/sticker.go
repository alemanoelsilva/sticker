package sticker

import (
	"errors"
	entity "sticker/internal/app/entity"
	model "sticker/internal/app/repository/mysql/model"

	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

func handleSqlError() error {
	return errors.New("internal Server Error")
}

type SqlRepository struct {
	DB     *gorm.DB
	Logger *zerolog.Logger
}

func (s *SqlRepository) AddSticker(details entity.Sticker, userId int) error {
	sticker := model.Sticker{
		ID:             details.ID,
		Name:           details.Name,
		Description:    details.Description,
		Category:       string(details.Category),
		Frequency:      string(details.Frequency),
		Status:         string(details.Status),
		IsPublic:       details.IsPublic,
		IsAutoApproval: details.IsAutoApproval,
		User: model.User{
			ID: userId,
		},
	}

	if result := s.DB.Create(&sticker); result.Error != nil {
		s.Logger.Error().Err(result.Error)
		return handleSqlError()
	}

	return nil
}

func (s *SqlRepository) UpdateStickerById(userId int, stickerId int, sticker entity.Sticker) error {
	var stickerModel model.Sticker

	if result := s.DB.Where(&model.Sticker{ID: stickerId, User: model.User{ID: userId}}).First(&stickerModel); result.Error != nil {
		s.Logger.Error().Err(result.Error)
		if result.Error.Error() == "sql: no rows in result set" {
			return errors.New("sticker not found")
		}
		return handleSqlError()
	}

	if stickerModel.ID == 0 {
		return errors.New("sticker not found")
	}

	stickerModel = model.Sticker{
		ID:             stickerId,
		Name:           sticker.Name,
		Description:    sticker.Description,
		Category:       string(sticker.Category),
		Frequency:      string(sticker.Frequency),
		Status:         string(sticker.Status),
		IsPublic:       sticker.IsPublic,
		IsAutoApproval: sticker.IsAutoApproval,
		UserId:         userId,
	}

	if result := s.DB.Save(&stickerModel); result.Error != nil {
		s.Logger.Error().Err(result.Error)
		return handleSqlError()
	}

	return nil
}

func (s *SqlRepository) GetStickerById(userId int, stickerId int) (detail entity.Sticker, err error) {
	stickerModel := model.Sticker{
		ID: stickerId,
		User: model.User{
			ID: userId,
		},
	}

	if result := s.DB.Where(&model.Sticker{ID: stickerId, User: model.User{ID: userId}}).First(&stickerModel); result.Error != nil {
		s.Logger.Error().Err(result.Error)
		if result.Error.Error() == "sql: no rows in result set" {
			return detail, errors.New("sticker not found")
		}
		return detail, handleSqlError()
	}

	detail = entity.Sticker{
		ID:             stickerModel.ID,
		Name:           stickerModel.Name,
		Description:    stickerModel.Description,
		Category:       entity.CategoryType(stickerModel.Category),
		Frequency:      entity.FrequencyType(stickerModel.Frequency),
		Status:         entity.StatusType(stickerModel.Status),
		IsPublic:       stickerModel.IsPublic,
		IsAutoApproval: stickerModel.IsAutoApproval,
	}

	return detail, nil
}

func (s *SqlRepository) GetStickers(userId int) (details []entity.Sticker, err error) {
	var stickerModel []model.Sticker

	if result := s.DB.Where(&model.Sticker{User: model.User{ID: userId}}).First(&stickerModel); result.Error != nil {
		s.Logger.Error().Err(result.Error)
		return details, handleSqlError()
	}

	details = make([]entity.Sticker, len(stickerModel))

	for i, s := range stickerModel {
		details[i] = entity.Sticker{
			ID:             s.ID,
			Name:           s.Name,
			Description:    s.Description,
			Category:       entity.CategoryType(s.Category),
			Frequency:      entity.FrequencyType(s.Frequency),
			Status:         entity.StatusType(s.Status),
			IsPublic:       s.IsPublic,
			IsAutoApproval: s.IsAutoApproval,
		}
	}

	return details, nil
}

func (s *SqlRepository) DeleteStickerById(userId int, stickerId int) error {
	stickerModel := model.Sticker{
		ID: stickerId,
	}
	result := s.DB.Where("user_id = ?", userId).Delete(&stickerModel)
	if result.Error != nil {
		s.Logger.Error().Err(result.Error)
		return handleSqlError()
	}

	if result.RowsAffected == 0 {
		return errors.New("sticker not found")
	}

	return nil
}
