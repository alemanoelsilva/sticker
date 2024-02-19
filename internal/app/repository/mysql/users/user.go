package user

import (
	"errors"
	"fmt"
	"sticker/internal/app/entity"
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

func (s *SqlRepository) AddUser(user entity.User) (err error) {
	userModel := model.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	if _, err = s.DB.NamedExec(query.AddUserQuery, &userModel); err != nil {
		s.Logger.Error().Err(err)
		return handleSqlError()
	}

	return nil
}

func (s *SqlRepository) GetUserById(id int) (user entity.User, err error) {
	var userModel model.User

	query := fmt.Sprintf(query.GetUserByIdQuery, id)

	if err = s.DB.Get(&userModel, query); err != nil {
		s.Logger.Error().Err(err)
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

	query := fmt.Sprintf(query.GetUserByEmailQuery, email)

	if err = s.DB.Get(&userModel, query); err != nil {
		s.Logger.Error().Err(err)
		return user, handleSqlError()
	}

	if userModel.ID == 0 {
		return user, errors.New("User not found")
	}

	user = entity.User{
		ID:       userModel.ID,
		Name:     userModel.Name,
		Email:    userModel.Email,
		Password: userModel.Password,
	}

	return user, nil
}
