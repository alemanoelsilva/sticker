package user

import (
	"errors"
	"sticker/internal/app/entity"
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

func (s *SqlRepository) AddUser(user entity.SignUp) error {
	userModel := model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	if result := s.DB.Save(&userModel); result.Error != nil {
		s.Logger.Error().Err(result.Error)
		return handleSqlError()
	}

	return nil
}

func (s *SqlRepository) GetUserById(id int) (user entity.User, err error) {
	var userModel model.User

	if result := s.DB.Where(&model.User{ID: id}).First(&userModel); result.Error != nil {
		s.Logger.Error().Err(result.Error)
		return user, handleSqlError()
	}

	user = entity.User{
		ID:       userModel.ID,
		Name:     userModel.Name,
		Email:    userModel.Email,
		Password: userModel.Password,
	}

	return user, nil
}

func (s *SqlRepository) GetUserByEmail(email string) (user entity.User, err error) {
	var userModel model.User

	if result := s.DB.Where(&model.User{Email: email}).First(&userModel); result.Error != nil {
		s.Logger.Error().Err(result.Error)
		return user, handleSqlError()
	}

	if userModel.ID == 0 {
		return user, errors.New("user not found")
	}

	user = entity.User{
		ID:       userModel.ID,
		Name:     userModel.Name,
		Email:    userModel.Email,
		Password: userModel.Password,
	}

	return user, nil
}
