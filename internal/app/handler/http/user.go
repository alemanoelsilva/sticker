package http

import (
	"net/http"
	"sticker/internal/app/entity"
	"sticker/internal/app/handler/http/validators"

	"github.com/gin-gonic/gin"
)

func (h *GinHandler) signUp(c *gin.Context) {
	response := ResponseJSON{c: c}

	var input entity.User

	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorHandler(http.StatusBadRequest, err)
		return
	}

	if err := h.Validator.Struct(validators.SignUp{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	}); err != nil {
		response.ErrorHandler(http.StatusBadRequest, err)
		return
	}

	err := h.UserUseCase.SignUp(input)
	if err != nil {
		response.ErrorHandler(http.StatusBadRequest, err)
		return
	}

	response.SuccessHandler(http.StatusCreated, map[string]interface{}{"message": "User created"})
}

func (h *GinHandler) signIn(c *gin.Context) {
	response := ResponseJSON{c: c}

	var input entity.SignIn

	if err := c.ShouldBindJSON(&input); err != nil {
		response.ErrorHandler(http.StatusBadRequest, err)
		return
	}

	if err := h.Validator.Struct(validators.SignIn{
		Email:    input.Email,
		Password: input.Password,
	}); err != nil {
		response.ErrorHandler(http.StatusBadRequest, err)
		return
	}

	token, err := h.UserUseCase.SignIn(input)
	if err != nil {
		response.ErrorHandler(http.StatusBadRequest, err)
		return
	}

	response.SuccessHandler(http.StatusOK, map[string]interface{}{"token": token})
}
