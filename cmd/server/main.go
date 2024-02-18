package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sticker/config"
	"sticker/config/database"

	handler "sticker/internal/app/handler/http"
	repo "sticker/internal/app/repository/mysql"

	stickerUseCase "sticker/internal/app/useCase/stickers"
	userUseCase "sticker/internal/app/useCase/users"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
)

func main() {
	// Set up ZeroLog logger
	logger := zerolog.New(os.Stdout)

	config.LoadAppConfig(&logger)

	ddb := database.DDB{Logger: &logger}
	db := ddb.Connect(config.AppConfig.ConnectionString)

	// TODO: check migration
	// ddb.RunMigrations()

	logger.Info().Msg("Initializing Repository (MySQL)")
	userRepository, stickerRepository := repo.NewSqlRepository(db)

	logger.Info().Msg("Initializing UseCases")
	userUseCaseService := userUseCase.LoadService(*userRepository, &logger)
	stickerUseCaseService := stickerUseCase.LoadService(*stickerRepository, &logger)

	logger.Info().Msg("Initializing Handlers")
	router := handler.NewGinHandler(*userUseCaseService, *stickerUseCaseService)

	// Start the server
	logger.Info().Msg(fmt.Sprintf("Starting Server on port %s", config.AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.AppConfig.Port), router))
}
