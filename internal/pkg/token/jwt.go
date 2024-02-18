package token

import (
	"errors"
	"sticker/config"
	"sticker/internal/app/entity"
	"time"

	"github.com/golang-jwt/jwt"
)

func NewAccessToken(user entity.SignIn) (string, error) {
	claims := entity.JwtClaims{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(config.AppConfig.MinutesToJwtExpire)).Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return accessToken.SignedString([]byte(config.AppConfig.SecretToken))
}

func ParseAccessToken(accessToken string) (*entity.JwtClaims, error) {
	parsedAccessToken, err := jwt.ParseWithClaims(accessToken, &entity.JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfig.SecretToken), nil
	})
	if err != nil {
		return nil, errors.New("Authorization token is invalid")
	}

	return parsedAccessToken.Claims.(*entity.JwtClaims), nil
}
