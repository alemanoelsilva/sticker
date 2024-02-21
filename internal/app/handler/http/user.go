package http

import (
	"net/http"
	"sticker/internal/app/entity"

	"github.com/labstack/echo/v4"
)

func LoadUserRoutes(router *echo.Echo, handler *EchoHandler) {
	router.POST("/api/v1/sign-up", handler.signUpHandler)
	router.POST("/api/v1/sign-in", handler.signInHandler)
}

func (e *EchoHandler) signUpHandler(c echo.Context) error {
	response := ResponseJSON{c: c}

	var input entity.User

	if err := c.Bind(&input); err != nil {
		return response.ErrorHandler(http.StatusBadRequest, err)
	}

	if err := e.UserUseCase.SignUp(input); err != nil {
		return response.ErrorHandler(http.StatusBadRequest, err)
	}

	return response.SuccessHandler(http.StatusCreated, handleResponseMessage("Sign up done"))
}

func (e *EchoHandler) signInHandler(c echo.Context) error {
	response := ResponseJSON{c: c}

	var input entity.SignIn

	if err := c.Bind(&input); err != nil {
		return response.ErrorHandler(http.StatusBadRequest, err)
	}

	token, err := e.UserUseCase.SignIn(input)
	if err != nil {
		return response.ErrorHandler(http.StatusBadRequest, err)
	}

	return response.SuccessHandler(http.StatusOK, map[string]interface{}{"token": token})
}
