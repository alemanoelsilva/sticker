package http

import (
	"net/http"
	"sticker/internal/app/entity"

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

	err := h.StickerUseCase.CreateSticker(input)
	if err != nil {
		r.ErrorHandler(http.StatusBadRequest, err)
		return
	}

	r.SuccessHandler(http.StatusOK, map[string]interface{}{"Message": "Sticker Created"})

}
