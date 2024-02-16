package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"

	"sticker/src/config"
	users_entity "sticker/src/entities/users"
)

func NewAccessToken(user users_entity.UserSignin) (string, error) {
	userClaims := users_entity.UserClaims{
		Email:    user.Email,
		Password: user.Password,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims)

	return accessToken.SignedString([]byte(config.AppConfig.SecretToken))
}

func NewRefreshToken(claims jwt.StandardClaims) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return refreshToken.SignedString([]byte(config.AppConfig.SecretToken))
}

func ParseAccessToken(accessToken string) (*users_entity.UserClaims, error) {
	parsedAccessToken, err := jwt.ParseWithClaims(accessToken, &users_entity.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.SecretToken), nil
	})
	if err != nil {
		return nil, errors.New("authorization token is invalid")
	}

	return parsedAccessToken.Claims.(*users_entity.UserClaims), nil
}

func ParseRefreshToken(refreshToken string) (*jwt.StandardClaims, error) {
	parsedRefreshToken, err := jwt.ParseWithClaims(refreshToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.SecretToken), nil
	})
	if err != nil {
		return nil, errors.New("authorization token is invalid")
	}

	return parsedRefreshToken.Claims.(*jwt.StandardClaims), nil
}
