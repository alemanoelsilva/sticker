package stickers

import (
	entity "sticker/internal/app/entity"
	validators "sticker/internal/app/handler/http/validators"
	repo "sticker/internal/app/repository/mysql/stickers"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
)

type Service struct {
	Validator  *validator.Validate
	Repository repo.SqlRepository
	Logger     *zerolog.Logger
}

func LoadService(v *validator.Validate, r repo.SqlRepository, l *zerolog.Logger) *Service {
	return &Service{
		Validator:  v,
		Logger:     l,
		Repository: r,
	}
}
func (ser *Service) CreateSticker(input entity.Sticker, userId int) error {
	ser.Logger.Info().Msg("Creating a Sticker")

	if err := ser.Validator.Struct(validators.Sticker{
		Name:           input.Name,
		Description:    input.Description,
		Category:       string(input.Category),
		Frequency:      string(input.Frequency),
		Status:         string(input.Status),
		IsPublic:       strconv.FormatBool(input.IsPublic),
		IsAutoApproval: strconv.FormatBool(input.IsAutoApproval),
	}); err != nil {
		return err
	}

	return ser.Repository.AddSticker(input, userId)
}

func (ser *Service) GetStickers(userId int) ([]entity.Sticker, error) {
	ser.Logger.Info().Msg("Getting Stickers")

	return ser.Repository.GetStickers(userId)
}

func (ser *Service) GetStickerById(userId int, stickerId int) (entity.Sticker, error) {
	ser.Logger.Info().Msg("Getting Sticker by id")

	return ser.Repository.GetStickerById(userId, stickerId)
}

func (ser *Service) UpdateStickerById(input entity.Sticker, userId int, stickerId int) error {
	ser.Logger.Info().Msg("Updating Sticker by id")

	if err := ser.Validator.Struct(validators.Sticker{
		Name:           input.Name,
		Description:    input.Description,
		Category:       string(input.Category),
		Frequency:      string(input.Frequency),
		Status:         string(input.Status),
		IsPublic:       strconv.FormatBool(input.IsPublic),
		IsAutoApproval: strconv.FormatBool(input.IsAutoApproval),
	}); err != nil {
		return err
	}

	return ser.Repository.UpdateStickerById(userId, stickerId, input)
}

func (ser *Service) DeleteStickerById(userId int, stickerId int) error {
	ser.Logger.Info().Msg("Deleting Sticker by id")

	return ser.Repository.DeleteStickerById(userId, stickerId)
}

func (ser *Service) InactivateStickerById(userId int, stickerId int) error {
	ser.Logger.Info().Msg("Inactivating Sticker by id")

	return ser.Repository.InactivateStickerById(userId, stickerId)
}
