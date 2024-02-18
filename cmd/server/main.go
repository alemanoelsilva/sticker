package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sticker/config"

	handler "sticker/internal/app/handler/http"
	mysqldb "sticker/internal/app/repository/mysql"
	useCase "sticker/internal/app/useCase"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
)

func main() {
	// Set up ZeroLog logger
	logger := zerolog.New(os.Stdout)

	config.LoadAppConfig(&logger)

	db, err := sqlx.Connect("mysql", config.AppConfig.ConnectionString)
	if err != nil {
		logger.Fatal().Err(err).Msg("MySQL connection error")
		os.Exit(1)
	}

	logger.Info().Msg("Running Migrations")
	mysqldb.RunMigrations(db, &logger)

	logger.Info().Msg("Initializing Repository (MySQL)")
	repo := mysqldb.NewMysqlDB(db)

	logger.Info().Msg("Initializing UseCases")
	useCaseService := useCase.LoadService(*repo, &logger)

	logger.Info().Msg("Initializing Handlers")
	router := handler.NewGinHandler(*useCaseService)

	// Start the server
	logger.Info().Msg(fmt.Sprintf("Starting Server on port %d", config.AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", config.AppConfig.Port), router))
}
