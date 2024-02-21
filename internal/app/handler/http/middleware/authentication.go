package middleware

import (
	"fmt"
	"net/http"
	"strings"

	jwt "sticker/internal/pkg/token"

	"github.com/labstack/echo/v4"
)

func TokenAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {

	return func(c echo.Context) error {
		const prefix = "Bearer "

		header := c.Request().Header.Get("Authorization")
		if header == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Authorization header missing or empty"})
		}

		token := strings.TrimPrefix(header, prefix)
		claims, err := jwt.ParseAccessToken(token)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{"error": fmt.Sprintf("unable to verify token with error: %v", err)})
		}

		c.Set("userId", claims.ID)

		return next(c)
	}
}
