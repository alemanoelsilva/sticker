package users

import (
	"errors"

	entity "sticker/internal/app/entity"
	validators "sticker/internal/app/handler/http/validators"
	repo "sticker/internal/app/repository/mysql/users"

	// TODO: implement interfaces to encrypt and jwt
	encrypt "sticker/internal/pkg/encrypt"
	jwt "sticker/internal/pkg/token"

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
		Repository: r,
		Logger:     l,
	}
}

func (ser *Service) SignUp(input entity.User) (err error) {
	ser.Logger.Info().Msg("Creating a User")

	if err = ser.Validator.Struct(validators.SignUp{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}); err != nil {
		return err
	}

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

	if err := ser.Validator.Struct(validators.SignIn{
		Email:    input.Email,
		Password: input.Password,
	}); err != nil {
		return "", err
	}

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

	input.ID = userModel.ID

	token, err = jwt.NewAccessToken(input)
	if err != nil {
		ser.Logger.Error().Msg(err.Error())
		return "", errors.New("An Internal error happened, contact the system admin")
	}

	return token, nil
}
