package http

import (
	stickerUseCase "sticker/internal/app/useCase/stickers"
	userUseCase "sticker/internal/app/useCase/users"

	"github.com/labstack/echo/v4"
)

type EchoHandler struct {
	UserUseCase    userUseCase.Service
	StickerUseCase stickerUseCase.Service
}

func NewEchoHandler(userUseCase userUseCase.Service, stickerUseCase stickerUseCase.Service) *echo.Echo {
	handler := &EchoHandler{
		UserUseCase:    userUseCase,
		StickerUseCase: stickerUseCase,
	}

	router := echo.New()

	// Users
	LoadUserRoutes(router, handler)
	// Stickers
	LoadStickerRoutes(router, handler)

	return router
}

func handleResponseMessage(msg string) interface{} {
	return map[string]interface{}{"message": msg}
}

type ResponseJSON struct {
	c echo.Context
}

func (s ResponseJSON) SuccessHandler(code int, data interface{}) error {
	return s.c.JSON(code, data)
}

func (s ResponseJSON) ErrorHandler(code int, err error) error {
	return s.c.JSON(code, map[string]interface{}{"error": err.Error()})
}
