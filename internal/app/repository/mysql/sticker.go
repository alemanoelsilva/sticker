package mysql

import (
	"fmt"
	"sticker/internal/app/entity"
	model "sticker/internal/app/repository/mysql/model"
	query "sticker/internal/app/repository/mysql/query"
)

func (s *MysqlDB) AddSticker(details entity.Sticker) (err error) {
	sticker := model.Sticker{
		ID:             details.ID,
		Name:           details.Name,
		Description:    details.Description,
		Category:       string(details.Category),
		Frequency:      string(details.Frequency),
		Status:         string(details.Status),
		IsPublic:       details.IsPublic,
		IsAutoApproved: details.IsAutoApproved,
	}

	query := fmt.Sprintf(query.AddStickerQuery)
	_, err = s.db.NamedExec(query, &sticker)
	return err
}

func (s *MysqlDB) UpdateStickerById(id int, details entity.Sticker) (err error) {
	sticker := model.Sticker{
		ID:             details.ID,
		Name:           details.Name,
		Description:    details.Description,
		Category:       string(details.Category),
		Frequency:      string(details.Frequency),
		Status:         string(details.Status),
		IsPublic:       details.IsPublic,
		IsAutoApproved: details.IsAutoApproved,
	}

	query := fmt.Sprintf(query.UpdateStickerByIdQuery, id)
	_, err = s.db.NamedExec(query, &sticker)
	return err
}

func (s *MysqlDB) GetStickerById(id int) (detail entity.Sticker, err error) {
	var stickerModel model.Sticker

	query := fmt.Sprintf(query.GetStickerByIdQuery, id)

	err = s.db.Get(&stickerModel, query)
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
		IsAutoApproved: stickerModel.IsAutoApproved,
	}

	return detail, nil
}

func (s *MysqlDB) GetStickers() (detail []entity.Sticker, err error) {
	var stickerModel []model.Sticker

	query := fmt.Sprintf(query.GetStickersQuery)

	err = s.db.Select(&stickerModel, query)
	if err != nil {
		return detail, err
	}

	for i, s := range stickerModel {
		detail[i].ID = s.ID
		detail[i].Description = s.Description
		detail[i].Category = entity.CategoryType(s.Category)
		detail[i].Frequency = entity.FrequencyType(s.Frequency)
		detail[i].Status = entity.StatusType(s.Status)
		detail[i].IsAutoApproved = s.IsAutoApproved
		detail[i].IsAutoApproved = s.IsAutoApproved
	}

	return detail, nil
}

func (s *MysqlDB) DeleteStickerById(id int) (err error) {
	query := fmt.Sprintf(query.DeleteStickerByIdQuery, id)
	_, err = s.db.Exec(query)

	if err != nil {
		return err
	}

	return nil
}
