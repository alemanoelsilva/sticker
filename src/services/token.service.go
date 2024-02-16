package services

import (
	"time"

	"github.com/golang-jwt/jwt"

	"sticker/src/config"
	"sticker/src/entities"
)

func NewAccessToken(user entities.User) (string, error) {
	userClaims := entities.UserClaims{
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

func ParseAccessToken(accessToken string) *entities.UserClaims {
	parsedAccessToken, _ := jwt.ParseWithClaims(accessToken, &entities.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.SecretToken), nil
	})

	return parsedAccessToken.Claims.(*entities.UserClaims)
}

func ParseRefreshToken(refreshToken string) *jwt.StandardClaims {
	parsedRefreshToken, _ := jwt.ParseWithClaims(refreshToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.SecretToken), nil
	})

	return parsedRefreshToken.Claims.(*jwt.StandardClaims)
}
