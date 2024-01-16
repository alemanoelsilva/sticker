package controllers

import (
	"encoding/json"
	"net/http"
	"sticker/src/entities"
	"sticker/src/services"

	"github.com/gorilla/mux"
)

func setHttpHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func CreateSticker(w http.ResponseWriter, r *http.Request) {
	setHttpHeaders(w)

	var sticker entities.Sticker
	json.NewDecoder(r.Body).Decode(&sticker)

	services.CreateSticker(&sticker)

	json.NewEncoder(w).Encode(sticker)
}

func GetStickers(w http.ResponseWriter, r *http.Request) {
	setHttpHeaders(w)

	stickers := services.GetStickers()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stickers)
}

func GetStickerById(w http.ResponseWriter, r *http.Request) {
	setHttpHeaders(w)
	stickerId := mux.Vars(r)["id"]

	sticker, err := services.GetStickerById(stickerId)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(sticker)
}
