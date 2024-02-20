package http

import (
	"errors"
	"net/http"
	"sticker/internal/app/entity"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getUserFromHeaders(c *gin.Context) int {
	userId := c.GetInt("userId")

	if userId == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "You are not logged in"})
		return 0
	}

	return userId
}

func getIdFromParams(c *gin.Context) (int, error) {
	idStr := c.Param("id")
	if idStr == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Sticker id is missing"})
		// TODO: return and handle it on caller level?
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Sticker id is not a number"})
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

	userId := getUserFromHeaders(c)

	if err := h.StickerUseCase.CreateSticker(input, userId); err != nil {
		response.ErrorHandler(http.StatusBadRequest, err)
		return
	}

	response.SuccessHandler(http.StatusCreated, handleResponseMessage("Sticker Created"))
}

func (h *GinHandler) getStickers(c *gin.Context) {
	response := ResponseJSON{c: c}

	userId := getUserFromHeaders(c)

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

	userId := getUserFromHeaders(c)

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

	userId := getUserFromHeaders(c)

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

	userId := getUserFromHeaders(c)

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
