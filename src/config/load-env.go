package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type EnvConfig struct {
	Port             int
	ConnectionString string
	SecretToken      string
}

var AppConfig *EnvConfig

func LoadAppConfig() {
	log.Println("Loading Server Configurations...")

	envPort := os.Getenv("PORT")
	envDBString := os.Getenv("DB_STRING")
	envJwtToken := os.Getenv("JWT_TOKEN")

	portInt, err := strconv.Atoi(envPort)
	if err != nil {
		fmt.Printf("%v\n", err)
		log.Fatalf("Error setting up the Application Port %v\n", envPort)
	}

	AppConfig = &EnvConfig{
		Port:             portInt,
		ConnectionString: envDBString,
		SecretToken:      envJwtToken,
	}

}
