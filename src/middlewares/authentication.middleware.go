package middlewares

import (
	"errors"
	"sticker/src/services"
	"time"
)

func IsAuthenticated(token string) error {
	if len(token) == 0 {
		return errors.New("authorization is required")
	}

	userClaims, err := services.ParseAccessToken(token)

	if err != nil || userClaims.VerifyExpiresAt(time.Now().Unix(), true) {
		return errors.New("user is not authenticated")

	}

	return nil
}
