package controllers

import (
	"encoding/json"
	"net/http"
	"sticker/src/entities"
	"sticker/src/services"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	setHttpHeaders(w)

	var user entities.User
	json.NewDecoder(r.Body).Decode(&user)

	// TODO!: validate user against DB

	token, err := services.NewAccessToken(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(token)
}
