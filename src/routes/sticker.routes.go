package routes

import (
	"sticker/src/controllers"

	"github.com/gorilla/mux"
)

// Initialize the router
var Router = mux.NewRouter().StrictSlash(true)

func RegisterStickerRoutes() {
	Router.HandleFunc("/api/stickers", controllers.GetStickers).Methods("GET")
	Router.HandleFunc("/api/stickers/{id}", controllers.GetStickerById).Methods("GET")
	Router.HandleFunc("/api/stickers", controllers.CreateSticker).Methods("POST")
	// Router.HandleFunc("/api/stickers/{id}", controllers.UpdateSticker).Methods("PUT")
	// Router.HandleFunc("/api/stickers/{id}", controllers.DeleteSticker).Methods("DELETE")
}
