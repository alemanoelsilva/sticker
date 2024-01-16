package main

import (
	"fmt"
	"log"
	"net/http"
	"sticker/src/config"
	"sticker/src/config/database"
	"sticker/src/routes"
)

func main() {
	// Load Configurations from config.json using Viper
	config.LoadAppConfig()

	// Initialize Database
	database.Connect(config.AppConfig.ConnectionString)
	database.Migrate()

	// Register Routes
	routes.RegisterStickerRoutes()

	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", config.AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.AppConfig.Port), routes.Router))
}
