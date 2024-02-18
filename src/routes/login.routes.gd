package routes

import (
	"sticker/src/controllers"

	"github.com/gorilla/mux"
)

func RegisterSignupRoutes(router *mux.Router) {
	router.HandleFunc("/api/signup", controllers.Signup).Methods("POST")
	router.HandleFunc("/api/signin", controllers.Signin).Methods("POST")
}
