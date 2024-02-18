package mysql

import (
	"fmt"
	"sticker/internal/app/entity"
	model "sticker/internal/app/repository/mysql/model"
	query "sticker/internal/app/repository/mysql/query"
)

func (s *MysqlDB) AddUser(user entity.User) (err error) {
	// TODO: add unique constraint to email
	userModel := model.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	_, err = s.db.NamedExec(query.AddUserQuery, &userModel)
	return err
}

func (s *MysqlDB) GetUserById(id int) (detail entity.User, err error) {
	var userModel model.User

	query := fmt.Sprintf(query.GetUserByIdQuery, id)

	err = s.db.Get(&userModel, query)
	if err != nil {
		return detail, err
	}

	detail = entity.User{
		ID:       userModel.ID,
		Name:     userModel.Name,
		Email:    userModel.Email,
		Password: userModel.Password,
	}

	return detail, nil
}

func (s *MysqlDB) GetUserByEmail(email string) (user entity.User, err error) {
	var userModel model.User

	query := fmt.Sprintf(query.GetUserByEmailQuery, email)

	err = s.db.Get(&userModel, query)
	if err != nil || userModel.ID == 0 {
		return user, err
	}

	user = entity.User{
		ID:       userModel.ID,
		Name:     userModel.Name,
		Email:    userModel.Email,
		Password: userModel.Password,
	}

	return user, nil
}
