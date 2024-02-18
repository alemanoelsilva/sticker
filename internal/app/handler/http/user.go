package http

import (
	"net/http"
	"sticker/internal/app/entity"

	"github.com/gin-gonic/gin"
)

func (h *GinHandler) signUp(c *gin.Context) {
	response := ResponseJSON{c: c}

	var input entity.User

	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorHandler(http.StatusBadRequest, err)
		return
	}

	// TODO: validate body request

	err := h.UserUseCase.SignUp(input)
	if err != nil {
		response.ErrorHandler(http.StatusBadRequest, err)
		return
	}

	response.SuccessHandler(http.StatusOK, map[string]interface{}{"Message": "User Created"})
}

func (h *GinHandler) signIn(c *gin.Context) {
	response := ResponseJSON{c: c}

	var input entity.SignIn

	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorHandler(http.StatusBadRequest, err)
		return
	}

	// TODO: validate body request

	token, err := h.UserUseCase.SignIn(input)
	if err != nil {
		response.ErrorHandler(http.StatusBadRequest, err)
		return
	}

	response.SuccessHandler(http.StatusOK, map[string]interface{}{"token": token})
}