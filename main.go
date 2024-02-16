package main

import (
	"fmt"
	"log"
	"net/http"
	"sticker/src/config"
	"sticker/src/config/database"
	"sticker/src/routes"

	"github.com/gorilla/mux"
)

func main() {
	// Load Configurations from config.json using Viper
	config.LoadAppConfig()

	// Initialize Database
	database.Connect(config.AppConfig.ConnectionString)
	database.Migrate()

	// Initialize the router
	var Router = mux.NewRouter().StrictSlash(true)

	// Register Routes
	routes.RegisterStickerRoutes(Router)
	routes.RegisterSignupRoutes(Router)

	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %d", config.AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.AppConfig.Port), Router))
}
