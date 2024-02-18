package token

import (
	"sticker/config"
	"sticker/internal/app/entity"
	"time"

	"github.com/golang-jwt/jwt"
)

func NewAccessToken(user entity.SignIn) (string, error) {
	// TODO: add user_id to claims
	claims := entity.JwtClaims{
		Email:    user.Email,
		Password: user.Password,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return accessToken.SignedString([]byte(config.AppConfig.SecretToken))
}
