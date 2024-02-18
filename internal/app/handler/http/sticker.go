package http

import (
	"errors"
	"net/http"
	"sticker/internal/app/entity"
	"sticker/internal/app/handler/http/validators"
	"sticker/internal/pkg/token"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func getUserFromToken(c *gin.Context) (userId int, err error, code int) {
	tokenString := c.Request.Header["Authorization"]
	if tokenString == nil {
		return 0, errors.New("Authentication is missing"), http.StatusBadRequest
	}

	// Split the token string by whitespace to separate "Bearer" from the token
	parts := strings.Fields(tokenString[0])

	// Check if there are two parts (Bearer and token)
	if len(parts) != 2 {
		return 0, errors.New("Authentication is missing"), http.StatusBadRequest
		// response.ErrorHandler(http.StatusBadRequest, errors.New("Authentication is missing"))
		// return
	}

	// Extract the token part
	authToken := parts[1]

	claims, err := token.ParseAccessToken(authToken)
	if err != err {
		return 0, err, http.StatusNonAuthoritativeInfo
		// response.ErrorHandler(http.StatusNonAuthoritativeInfo, err)
		// return
	}

	return claims.ID, nil, 0
}

func (h *GinHandler) createSticker(c *gin.Context) {
	response := ResponseJSON{c: c}

	var input entity.Sticker

	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorHandler(http.StatusBadRequest, err)
		return
	}

	userId, err, code := getUserFromToken(c)
	if err != nil {
		response.ErrorHandler(code, err)
		return
	}

	if err := h.Validator.Struct(validators.Sticker{
		Name:           input.Name,
		Description:    input.Description,
		Category:       string(input.Category),
		Frequency:      string(input.Frequency),
		Status:         string(input.Status),
		IsPublic:       strconv.FormatBool(input.IsPublic),
		IsAutoApproval: strconv.FormatBool(input.IsAutoApproval),
	}); err != nil {
		response.ErrorHandler(http.StatusBadRequest, err)
		return
	}

	if err := h.StickerUseCase.CreateSticker(input, userId); err != nil {
		response.ErrorHandler(http.StatusBadRequest, err)
		return
	}

	response.SuccessHandler(http.StatusCreated, map[string]interface{}{"message": "Sticker Created"})
}

func (h *GinHandler) getStickers(c *gin.Context) {
	response := ResponseJSON{c: c}

	userId, err, code := getUserFromToken(c)
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

	userId, err, code := getUserFromToken(c)
	if err != nil {
		response.ErrorHandler(code, err)
		return
	}

	stickerIdStr := c.Param("id")
	if stickerIdStr == "" {
		response.ErrorHandler(http.StatusBadRequest, errors.New("Sticker id is missing"))
	}

	stickerId, err := strconv.Atoi(stickerIdStr)
	if err != nil {
		response.ErrorHandler(http.StatusBadRequest, errors.New("Sticker id is not a number"))
		return
	}

	stickers, err := h.StickerUseCase.GetStickerById(userId, stickerId)
	if err != nil {
		response.ErrorHandler(http.StatusBadRequest, err)
		return
	}

	response.SuccessHandler(http.StatusOK, stickers)
}

func (h *GinHandler) updateStickerById(c *gin.Context) {
	response := ResponseJSON{c: c}

	var input entity.Sticker

	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorHandler(http.StatusBadRequest, err)
		return
	}

	userId, err, code := getUserFromToken(c)
	if err != nil {
		response.ErrorHandler(code, err)
		return
	}

	stickerIdStr := c.Param("id")
	if stickerIdStr == "" {
		response.ErrorHandler(http.StatusBadRequest, errors.New("Sticker id is missing"))
	}

	stickerId, err := strconv.Atoi(stickerIdStr)
	if err != nil {
		response.ErrorHandler(http.StatusBadRequest, errors.New("Sticker id is not a number"))
		return
	}

	if err := h.Validator.Struct(validators.Sticker{
		Name:           input.Name,
		Description:    input.Description,
		Category:       string(input.Category),
		Frequency:      string(input.Frequency),
		Status:         string(input.Status),
		IsPublic:       strconv.FormatBool(input.IsPublic),
		IsAutoApproval: strconv.FormatBool(input.IsAutoApproval),
	}); err != nil {
		response.ErrorHandler(http.StatusBadRequest, err)
		return
	}

	if err := h.StickerUseCase.UpdateStickerById(input, userId, stickerId); err != nil {
		response.ErrorHandler(http.StatusBadRequest, err)
		return
	}

	response.SuccessHandler(http.StatusOK, map[string]interface{}{"message": "Sticker Updated"})
}
