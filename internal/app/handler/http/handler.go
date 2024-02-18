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

	// Stickers
	router.POST("/api/v1/stickers", handler.createSticker)
	router.GET("/api/v1/stickers", handler.getStickers)

	// Users
	router.POST("/api/v1/sign-up", handler.signUp)
	router.POST("/api/v1/sign-in", handler.signIn)

	return router
}

func now() time.Time {
	return time.Now()
}

type ResponseJSON struct {
	c *gin.Context
}

func (s ResponseJSON) SuccessHandler(code int, data interface{}) {
	s.c.JSON(code, data)
}

func (s ResponseJSON) ErrorHandler(code int, err error) {
	s.c.JSON(code, gin.H{"error": err.Error()})
}
