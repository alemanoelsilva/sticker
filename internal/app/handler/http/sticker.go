package http

import (
	"errors"
	"net/http"
	"sticker/internal/app/entity"
	"sticker/internal/pkg/token"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *GinHandler) createSticker(c *gin.Context) {
	r := ResponseJSON{c: c}

	var input entity.Sticker

	if err := c.ShouldBindJSON(&input); err != nil {
		r.ErrorHandler(http.StatusBadRequest, err)
		return
	}

	// TODO: validate body request
	// TODO: validate token
	// TODO: extract user_id from token

	tokenString := c.Request.Header["Authorization"][0]
	// Split the token string by whitespace to separate "Bearer" from the token
	parts := strings.Fields(tokenString)

	// Check if there are two parts (Bearer and token)
	if len(parts) != 2 {
		r.ErrorHandler(http.StatusBadRequest, errors.New("Authentication is missing"))
		return
	}

	// Extract the token part
	authToken := parts[1]

	claims, err := token.ParseAccessToken(authToken)
	if err != err {
		r.ErrorHandler(http.StatusNonAuthoritativeInfo, err)
		return
	}

	if err := h.StickerUseCase.CreateSticker(input, claims.ID); err != nil {
		r.ErrorHandler(http.StatusBadRequest, err)
		return
	}

	r.SuccessHandler(http.StatusCreated, map[string]interface{}{"Message": "Sticker Created"})
}
	if err != nil {
		r.ErrorHandler(http.StatusBadRequest, err)
		return
	}

	r.SuccessHandler(http.StatusOK, map[string]interface{}{"Message": "Sticker Created"})

}
