package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvVars struct {
	PORT       string
	JWT_SECRET string
	GO_ENV     string
}

func LoadConfig() (EnvVars, error) {
	mode := os.Getenv("GO_ENV")

	if mode == "development" {
		log.Print("Running in development mode")
		err := godotenv.Load("./.development.env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	} else {
		log.Print("Running in production mode")
		err := godotenv.Load("./.production.env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	cfg := EnvVars{
		PORT:       os.Getenv("PORT"),
		JWT_SECRET: os.Getenv("JWT_SECRET"),
	}

	if cfg.PORT == "" {
		log.Fatal("PORT is not set")
	}

	if cfg.JWT_SECRET == "" {
		log.Fatal("JWT_SECRET is not set")
	}

	return cfg, nil
}
