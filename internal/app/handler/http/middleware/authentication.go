package middleware

import (
	"fmt"
	"net/http"

	jwt "sticker/internal/pkg/token"

	"github.com/gin-gonic/gin"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		const prefix = "Bearer "

		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Authorization header missing or empty"})
		}

		token := header[len(prefix):]
		claims, err := jwt.ParseAccessToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("unable to verify token with error: %v", err)})
		}

		c.Set("userId", claims.ID)

		c.Next()
	}
}
