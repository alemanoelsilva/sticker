package controllers

import (
	"encoding/json"
	"net/http"
	users_entity "sticker/src/entities/users"
	"sticker/src/services"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type SuccessResponse struct {
	Token string `json:"token"`
}

func Signup(w http.ResponseWriter, r *http.Request) {
	setHttpHeaders(w)

	var user users_entity.User
	json.NewDecoder(r.Body).Decode(&user)

	insertedUser := services.CreateUser(&user)

	userSignup := users_entity.UserSignin{
		Email:    insertedUser.Email,
		Password: insertedUser.Password,
	}

	token, err := services.NewAccessToken(userSignup)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(token)
}

func Signin(w http.ResponseWriter, r *http.Request) {
	setHttpHeaders(w)

	var userSignin users_entity.UserSignin
	json.NewDecoder(r.Body).Decode(&userSignin)

	err := services.ValidateUserByEmailAndPassword(userSignin.Email, userSignin.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
		return
	}

	token, err := services.NewAccessToken(userSignin)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Message: err.Error()})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(SuccessResponse{Token: token})
}
