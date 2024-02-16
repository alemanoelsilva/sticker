package routes

import (
	"sticker/src/controllers"

	"github.com/gorilla/mux"
)

func RegisterStickerRoutes(router *mux.Router) {
	router.HandleFunc("/api/stickers", controllers.GetStickers).Methods("GET")
	router.HandleFunc("/api/stickers/{id}", controllers.GetStickerById).Methods("GET")
	router.HandleFunc("/api/stickers", controllers.CreateSticker).Methods("POST")
	router.HandleFunc("/api/stickers/{id}", controllers.UpdateSticker).Methods("PUT")
	router.HandleFunc("/api/stickers/{id}", controllers.DeleteStickerById).Methods("DELETE")
}
