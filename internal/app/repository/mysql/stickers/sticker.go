package sticker

import (
	"errors"
	"fmt"
	entity "sticker/internal/app/entity"
	model "sticker/internal/app/repository/mysql/model"
	query "sticker/internal/app/repository/mysql/query"

	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

func handleSqlError() error {
	return errors.New("Internal Server Error")
}

type SqlRepository struct {
	DB     *sqlx.DB
	Logger *zerolog.Logger
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
	if _, err = s.DB.NamedExec(query, &sticker); err != nil {
		s.Logger.Error().Err(err)
		return handleSqlError()
	}
	return err
}

func (s *SqlRepository) UpdateStickerById(userId int, stickerId int, sticker entity.Sticker) (err error) {
	var stickerModel model.Sticker
	if err = s.DB.Get(&stickerModel, fmt.Sprintf(query.GetStickerByIdQuery, userId, stickerId)); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return errors.New("Sticker not found")
		}
		s.Logger.Error().Err(err)
		return handleSqlError()
	}

	if stickerModel.ID == 0 {
		return errors.New("Sticker not found")
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
	}

	query := fmt.Sprintf(query.UpdateStickerByIdQuery, userId, stickerId)
	if _, err = s.DB.NamedExec(query, &stickerModel); err != nil {
		s.Logger.Error().Err(err)
		return handleSqlError()
	}

	return nil
}

func (s *SqlRepository) GetStickerById(userId int, stickerId int) (detail entity.Sticker, err error) {
	var stickerModel model.Sticker

	query := fmt.Sprintf(query.GetStickerByIdQuery, userId, stickerId)

	if err = s.DB.Get(&stickerModel, query); err != nil {
		if err.Error() == "sql: no rows in result set" {
			return detail, errors.New("Sticker not found")
		}
		s.Logger.Error().Err(err)
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

	query := fmt.Sprintf(query.GetStickersQuery, userId)

	if err := s.DB.Select(&stickerModel, query); err != nil {
		s.Logger.Error().Err(err)
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

func (s *SqlRepository) DeleteStickerById(userId int, stickerId int) (err error) {
	query := fmt.Sprintf(query.DeleteStickerByIdQuery, userId, stickerId)

	result, err := s.DB.Exec(query)
	if err != nil {
		s.Logger.Error().Err(err)
		return handleSqlError()
	}

	affect, err := result.RowsAffected()
	if err != nil {
		s.Logger.Error().Err(err)
		return handleSqlError()
	}

	if affect == 0 {
		return errors.New("Sticker not found")
	}

	return nil
}
