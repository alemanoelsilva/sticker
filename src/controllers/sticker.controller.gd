package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	stickers_entity "sticker/src/entities/stickers"
	"sticker/src/middlewares"
	"sticker/src/services"
	"strconv"

	"github.com/gorilla/mux"
)

func setHttpHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func getIdFromPath(r *http.Request) (int, error) {
	idString := mux.Vars(r)["id"]

	stickerId, err := strconv.Atoi(idString)
	if err != nil {
		return 0, errors.New("Sticker id must be a number")
	}

	return stickerId, nil
}

func CreateSticker(w http.ResponseWriter, r *http.Request) {
	setHttpHeaders(w)
	auth := r.Header.Get("Authorization")
	err := middlewares.IsAuthenticated(auth)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	var sticker stickers_entity.Sticker
	json.NewDecoder(r.Body).Decode(&sticker)

	services.CreateSticker(&sticker)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(sticker)
}

func GetStickers(w http.ResponseWriter, r *http.Request) {
	setHttpHeaders(w)
	err := middlewares.IsAuthenticated(r.Header.Get("Authorization"))
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	stickers := services.GetStickers()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stickers)
}

func GetStickerById(w http.ResponseWriter, r *http.Request) {
	setHttpHeaders(w)
	err := middlewares.IsAuthenticated(r.Header.Get("Authorization"))
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	stickerId := mux.Vars(r)["id"]

	sticker, err := services.GetStickerById(stickerId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sticker)
}

func UpdateSticker(w http.ResponseWriter, r *http.Request) {
	setHttpHeaders(w)
	err := middlewares.IsAuthenticated(r.Header.Get("Authorization"))
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	stickerId, err := getIdFromPath(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	var sticker stickers_entity.Sticker
	json.NewDecoder(r.Body).Decode(&sticker)

	stickerUpdated, err := services.UpdateSticker(stickerId, &sticker)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stickerUpdated)
}

func DeleteStickerById(w http.ResponseWriter, r *http.Request) {
	setHttpHeaders(w)
	err := middlewares.IsAuthenticated(r.Header.Get("Authorization"))
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		json.NewEncoder(w).Encode(ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	stickerId, err := getIdFromPath(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	_, err = services.DeleteStickerById(stickerId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
