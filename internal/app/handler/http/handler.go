package http

import (
	stickerUseCase "sticker/internal/app/useCase/stickers"
	userUseCase "sticker/internal/app/useCase/users"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type GinHandler struct {
	UserUseCase    userUseCase.Service
	StickerUseCase stickerUseCase.Service
}

func NewGinHandler(userUseCase userUseCase.Service, stickerUseCase stickerUseCase.Service) *gin.Engine {
	handler := &GinHandler{
		UserUseCase:    userUseCase,
		StickerUseCase: stickerUseCase,
	}

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		// Start timer
		start := now()

		// Process request
		c.Next()

		// Log request details
		log.Info().
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Int("status", c.Writer.Status()).
			Dur("duration", now().Sub(start)).
			Msg("Request handled")
	})

	// Users
	LoadUserRoutes(router, handler)
	// Stickers
	LoadStickerRoutes(router, handler)

	return router
}

func now() time.Time {
	return time.Now()
}

func handleResponseMessage(msg string) interface{} {
	return map[string]interface{}{"message": msg}
}

type ResponseJSON struct {
	c *gin.Context
}

func (s ResponseJSON) SuccessHandler(code int, data interface{}) {
	s.c.IndentedJSON(code, data)
}

func (s ResponseJSON) ErrorHandler(code int, err error) {
	s.c.IndentedJSON(code, gin.H{"error": err.Error()})
}
