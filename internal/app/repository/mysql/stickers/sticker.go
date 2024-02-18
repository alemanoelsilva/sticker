package sticker

import (
	"errors"
	"fmt"
	"sticker/internal/app/entity"
	model "sticker/internal/app/repository/mysql/model"
	query "sticker/internal/app/repository/mysql/query"

	"github.com/jmoiron/sqlx"
)

type SqlRepository struct {
	DB *sqlx.DB
}

func (s *SqlRepository) AddSticker(details entity.Sticker, userId int) (err error) {
	sticker := model.Sticker{
		ID:             details.ID,
		Name:           details.Name,
		Description:    details.Description,
		Category:       string(details.Category),
		Frequency:      string(details.Frequency),
		Status:         string(details.Status),
		IsPublic:       details.IsPublic,
		IsAutoApproval: details.IsAutoApproval,
		UserId:         userId,
	}

	query := fmt.Sprintf(query.AddStickerQuery)
	_, err = s.DB.NamedExec(query, &sticker)
	return err
}

func (s *SqlRepository) UpdateStickerById(userId int, stickerId int, sticker entity.Sticker) (err error) {
	stickerModel := model.Sticker{
		ID:             stickerId,
		Name:           sticker.Name,
		Description:    sticker.Description,
		Category:       string(sticker.Category),
		Frequency:      string(sticker.Frequency),
		Status:         string(sticker.Status),
		IsPublic:       sticker.IsPublic,
		IsAutoApproval: sticker.IsAutoApproval,
	}

	query := fmt.Sprintf(query.UpdateStickerByIdQuery, userId, stickerId)
	_, err = s.DB.NamedExec(query, &stickerModel)
	return err
}

func (s *SqlRepository) GetStickerById(userId int, stickerId int) (detail entity.Sticker, err error) {
	var stickerModel model.Sticker

	query := fmt.Sprintf(query.GetStickerByIdQuery, userId, stickerId)

	err = s.DB.Get(&stickerModel, query)
	if err != nil {
		return detail, err
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

func (s *SqlRepository) GetStickers(userId int) ([]entity.Sticker, error) {
	var stickerModel []model.Sticker

	query := fmt.Sprintf(query.GetStickersQuery, userId)

	if err := s.DB.Select(&stickerModel, query); err != nil {
		return []entity.Sticker{}, err
	}

	details := make([]entity.Sticker, len(stickerModel))

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

func (s *SqlRepository) DeleteStickerById(userId int, stickerId int) (err error) {
	query := fmt.Sprintf(query.DeleteStickerByIdQuery, userId, stickerId)

	result, err := s.DB.Exec(query)
	if err != nil {
		return err
	}

	affect, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affect == 0 {
		return errors.New("Sticker not found")
	}

	return nil
}
