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

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
)

func main() {
	// Set up ZeroLog logger
	logger := zerolog.New(os.Stdout)
	validator := *validator.New()

	config.LoadAppConfig(&logger)

	ddb := database.DDB{Logger: &logger}
	db := ddb.Connect(config.AppConfig.ConnectionString)
	ddb.RunMigrations(db)

	logger.Info().Msg("Initializing Repository (MySQL)")
	// TODO: split repositories
	userRepository, stickerRepository := repo.NewSqlRepository(db, &logger)

	logger.Info().Msg("Initializing UseCases")
	userUseCaseService := userUseCase.LoadService(&validator, *userRepository, &logger)
	stickerUseCaseService := stickerUseCase.LoadService(&validator, *stickerRepository, &logger)

	logger.Info().Msg("Initializing Handlers")
	router := handler.NewEchoHandler(*userUseCaseService, *stickerUseCaseService)

	// Start the server
	logger.Info().Msg(fmt.Sprintf("Starting Server on port %s", config.AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.AppConfig.Port), router))
}
