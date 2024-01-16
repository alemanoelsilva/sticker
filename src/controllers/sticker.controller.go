package controllers

import (
	"encoding/json"
	"net/http"
	"sticker/src/config/database"
	"sticker/src/entities"
)

func CreateSticker(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var sticker entities.Sticker
	json.NewDecoder(r.Body).Decode(&sticker)
	database.Instance.Create(&sticker)
	json.NewEncoder(w).Encode(sticker)
}

func GetStickers(w http.ResponseWriter, r *http.Request) {
	var stickers []entities.Sticker
	database.Instance.Find(&stickers)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(stickers)
}
