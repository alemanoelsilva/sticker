package http

import (
	"errors"
	"net/http"
	"sticker/internal/app/entity"
	"sticker/internal/pkg/token"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func getUserFromToken(c *gin.Context) (int, int, error) {
	tokenString := c.Request.Header["Authorization"]
	if tokenString == nil {
		return 0, http.StatusBadRequest, errors.New("Authentication is missing")
	}

	// Split the token string by whitespace to separate "Bearer" from the token
	parts := strings.Fields(tokenString[0])

	// Check if there are two parts (Bearer and token)
	if len(parts) != 2 {
		return 0, http.StatusBadRequest, errors.New("Authentication is missing")
	}

	// Extract the token part
	authToken := parts[1]

	claims, err := token.ParseAccessToken(authToken)
	if err != err {
		return 0, http.StatusNonAuthoritativeInfo, err
	}

	return claims.ID, 0, nil
}

func getIdFromParams(c *gin.Context) (int, error) {
	idStr := c.Param("id")
	if idStr == "" {
		return 0, errors.New("Sticker id is missing")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return id, errors.New("Sticker id is not a number")
	}

	return id, nil
}

func (h *GinHandler) createSticker(c *gin.Context) {
	response := ResponseJSON{c: c}

	var input entity.Sticker

	if err := c.BindJSON(&input); err != nil {
		response.ErrorHandler(http.StatusBadRequest, err)
		return
	}

	userId, code, err := getUserFromToken(c)
	if err != nil {
		response.ErrorHandler(code, err)
		return
	}

	if err := h.StickerUseCase.CreateSticker(input, userId); err != nil {
		response.ErrorHandler(http.StatusBadRequest, err)
		return
	}

	response.SuccessHandler(http.StatusCreated, handleResponseMessage("Sticker Created"))
}

func (h *GinHandler) getStickers(c *gin.Context) {
	response := ResponseJSON{c: c}

	userId, code, err := getUserFromToken(c)
	if err != nil {
		response.ErrorHandler(code, err)
		return
	}

	stickers, err := h.StickerUseCase.GetStickers(userId)
	if err != nil {
		response.ErrorHandler(http.StatusBadRequest, err)
		return
	}

	response.SuccessHandler(http.StatusOK, stickers)
}

func (h *GinHandler) getStickerById(c *gin.Context) {
	response := ResponseJSON{c: c}

	tokenString := c.Request.Header["Authorization"]
	if tokenString == nil {
		response.ErrorHandler(http.StatusBadRequest, errors.New("Authentication is missing"))
		return
	}

	userId, code, err := getUserFromToken(c)
	if err != nil {
		response.ErrorHandler(code, err)
		return
	}

	id, err := getIdFromParams(c)
	if err != nil {
		response.ErrorHandler(http.StatusBadRequest, err)
		return
	}

	stickers, err := h.StickerUseCase.GetStickerById(userId, id)
	if err != nil {
		response.ErrorHandler(http.StatusBadRequest, err)
		return
	}

	response.SuccessHandler(http.StatusOK, stickers)
}

func (h *GinHandler) updateStickerById(c *gin.Context) {
	response := ResponseJSON{c: c}

	var input entity.Sticker

	if err := c.BindJSON(&input); err != nil {
		response.ErrorHandler(http.StatusBadRequest, err)
		return
	}

	userId, code, err := getUserFromToken(c)
	if err != nil {
		response.ErrorHandler(code, err)
		return
	}

	id, err := getIdFromParams(c)
	if err != nil {
		response.ErrorHandler(http.StatusBadRequest, err)
		return
	}

	if err := h.StickerUseCase.UpdateStickerById(input, userId, id); err != nil {
		response.ErrorHandler(http.StatusBadRequest, err)
		return
	}

	response.SuccessHandler(http.StatusOK, handleResponseMessage("Sticker Updated"))
}

func (h *GinHandler) deleteStickerById(c *gin.Context) {
	response := ResponseJSON{c: c}

	userId, code, err := getUserFromToken(c)
	if err != nil {
		response.ErrorHandler(code, err)
		return
	}

	id, err := getIdFromParams(c)
	if err != nil {
		response.ErrorHandler(http.StatusBadRequest, err)
		return
	}

	if err := h.StickerUseCase.DeleteStickerById(userId, id); err != nil {
		response.ErrorHandler(http.StatusBadRequest, err)
		return
	}

	response.SuccessHandler(http.StatusOK, handleResponseMessage("Sticker Removed"))
}
