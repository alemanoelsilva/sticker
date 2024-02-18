package users

import (
	"errors"

	"sticker/internal/app/entity"
	// TODO: implement interfaces to encrypt and jwt
	encrypt "sticker/internal/pkg/encrypt"
	jwt "sticker/internal/pkg/token"
)

func (ser *Service) SignUp(input entity.User) error {
	ser.Logger.Info().Msg("Creating a User")

	hashedPassword, err := encrypt.Hash(input.Password)
	if err != nil {
		ser.Logger.Error().Msg(err.Error())
		return errors.New("An Internal error happened, contact the system admin")
	}

	input.Password = hashedPassword

	return ser.Repository.AddUser(input)
}

func (ser *Service) SignIn(input entity.SignIn) (token string, err error) {
	ser.Logger.Info().Msg("Signing in a User")

	userModel, err := ser.Repository.GetUserByEmail(input.Email)
	if err != nil {
		ser.Logger.Error().Msg(err.Error())
		return "", errors.New("Credential does not match")
	}

	if userModel.ID == 0 {
		return "", errors.New("Credential not found")
	}

	if !encrypt.Check(input.Password, userModel.Password) {
		return "", errors.New("Wrong password")
	}

	token, err = jwt.NewAccessToken(input)
	if err != nil {
		ser.Logger.Error().Msg(err.Error())
		return "", errors.New("An Internal error happened, contact the system admin")
	}

	return token, nil
}
