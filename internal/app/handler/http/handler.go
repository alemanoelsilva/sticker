package http

import (
	"sticker/internal/app/useCase"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type GinHandler struct {
	UseCase useCase.Service
}

func NewGinHandler(useCase useCase.Service) *gin.Engine {
	handler := &GinHandler{
		UseCase: useCase,
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

	router.POST("/api/v1/stickers", handler.createSticker)
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
