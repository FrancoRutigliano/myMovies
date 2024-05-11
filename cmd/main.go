package main

import (
	"log"
	"os"

	"github.com/FrancoRutigliano/myMovies/cmd/api"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	portStr := ":" + os.Getenv("PORT")
	if portStr == "" {
		log.Fatal("PORT is not set in .env file")
	}

	app := api.NewAPIServer(portStr)

	if err := app.Run(); err != nil {
		log.Fatal("error to inicialize the server")
	}
}
