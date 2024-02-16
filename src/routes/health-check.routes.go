package routes

import (
	"sticker/src/controllers"

	"github.com/gorilla/mux"
)

func RegisterHealthCheckRoutes(router *mux.Router) {
	router.HandleFunc("/api/health-check", controllers.HealthCheck).Methods("GET")
}
