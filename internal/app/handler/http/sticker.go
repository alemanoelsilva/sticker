package http

import (
	"errors"
	"net/http"
	"sticker/internal/app/entity"
	middleware "sticker/internal/app/handler/http/middleware"
	"strconv"

	"github.com/labstack/echo/v4"
)

func getUserIdFromHeaders(c echo.Context) (int, error) {
	userId := c.Get("userId").(int)
	if userId == 0 {
		return 0, errors.New("you are not logged in")
	}

	return userId, nil
}

func getIdFromParams(c echo.Context) (int, error) {
	idStr := c.Param("id")
	if idStr == "" {
		return 0, errors.New("sticker id is missing")
		// TODO: return and handle it on caller level?
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, errors.New("sticker id is not a number")
	}

	return id, nil
}

func LoadStickerRoutes(router *echo.Echo, handler *EchoHandler) {
	router.POST("/api/v1/stickers", handler.createStickerHandler, middleware.TokenAuthMiddleware)
	router.GET("/api/v1/stickers", handler.getStickersHandler, middleware.TokenAuthMiddleware)
	router.GET("/api/v1/stickers/:id", handler.getStickerByIdHandler, middleware.TokenAuthMiddleware)
	router.PUT("/api/v1/stickers/:id", handler.updateStickerByIdHandler, middleware.TokenAuthMiddleware)
	router.DELETE("/api/v1/stickers/:id", handler.deleteStickerByIdHandler, middleware.TokenAuthMiddleware)
	router.PATCH("/api/v1/stickers/:id/inactivate", handler.inactivateStickerByIdHandler, middleware.TokenAuthMiddleware)
}

func (e *EchoHandler) createStickerHandler(c echo.Context) error {
	response := ResponseJSON{c: c}

	var input entity.Sticker

	if err := c.Bind(&input); err != nil {
		return response.ErrorHandler(http.StatusBadRequest, err)
	}

	userId, err := getUserIdFromHeaders(c)
	if err != nil {
		return response.ErrorHandler(http.StatusUnauthorized, err)
	}

	if err := e.StickerUseCase.CreateSticker(input, userId); err != nil {
		return response.ErrorHandler(http.StatusBadRequest, err)
	}

	return response.SuccessHandler(http.StatusCreated, handleResponseMessage("Sticker Created"))
}

func (e *EchoHandler) getStickersHandler(c echo.Context) error {
	response := ResponseJSON{c: c}

	userId, err := getUserIdFromHeaders(c)
	if err != nil {
		return response.ErrorHandler(http.StatusUnauthorized, err)
	}

	stickers, err := e.StickerUseCase.GetStickers(userId)
	if err != nil {
		return response.ErrorHandler(http.StatusBadRequest, err)
	}

	return response.SuccessHandler(http.StatusOK, stickers)
}

func (e *EchoHandler) getStickerByIdHandler(c echo.Context) error {
	response := ResponseJSON{c: c}

	userId, err := getUserIdFromHeaders(c)
	if err != nil {
		return response.ErrorHandler(http.StatusUnauthorized, err)
	}

	id, err := getIdFromParams(c)
	if err != nil {
		return response.ErrorHandler(http.StatusBadRequest, err)
	}

	stickers, err := e.StickerUseCase.GetStickerById(userId, id)
	if err != nil {
		return response.ErrorHandler(http.StatusBadRequest, err)
	}

	return response.SuccessHandler(http.StatusOK, stickers)
}

func (e *EchoHandler) updateStickerByIdHandler(c echo.Context) error {
	response := ResponseJSON{c: c}

	var input entity.Sticker

	if err := c.Bind(&input); err != nil {
		return response.ErrorHandler(http.StatusBadRequest, err)
	}

	userId, err := getUserIdFromHeaders(c)
	if err != nil {
		return response.ErrorHandler(http.StatusUnauthorized, err)
	}

	id, err := getIdFromParams(c)
	if err != nil {
		return response.ErrorHandler(http.StatusBadRequest, err)
	}

	if err := e.StickerUseCase.UpdateStickerById(input, userId, id); err != nil {
		return response.ErrorHandler(http.StatusBadRequest, err)
	}

	return response.SuccessHandler(http.StatusOK, handleResponseMessage("Sticker Updated"))
}

func (e *EchoHandler) deleteStickerByIdHandler(c echo.Context) error {
	response := ResponseJSON{c: c}

	userId, err := getUserIdFromHeaders(c)
	if err != nil {
		return response.ErrorHandler(http.StatusUnauthorized, err)
	}

	id, err := getIdFromParams(c)
	if err != nil {
		return response.ErrorHandler(http.StatusBadRequest, err)
	}

	if err := e.StickerUseCase.DeleteStickerById(userId, id); err != nil {
		return response.ErrorHandler(http.StatusBadRequest, err)
	}

	return response.SuccessHandler(http.StatusOK, handleResponseMessage("Sticker Removed"))
}

func (e *EchoHandler) inactivateStickerByIdHandler(c echo.Context) error {
	response := ResponseJSON{c: c}

	userId, err := getUserIdFromHeaders(c)
	if err != nil {
		return response.ErrorHandler(http.StatusUnauthorized, err)
	}

	stickerId, err := getIdFromParams(c)
	if err != nil {
		return response.ErrorHandler(http.StatusBadRequest, err)
	}

	if err := e.StickerUseCase.InactivateStickerById(userId, stickerId); err != nil {
		return response.ErrorHandler(http.StatusBadRequest, err)
	}

	return response.SuccessHandler(http.StatusOK, handleResponseMessage("Sticker inactivated"))
}
