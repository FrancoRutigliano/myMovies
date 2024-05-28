package main

import (
	"log"

	"github.com/FrancoRutigliano/myMovies/cmd/api"
	"github.com/FrancoRutigliano/myMovies/config"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("error to inicialize the server: %v", err)
	}

	port := ":" + config.PORT
	if port == "" {
		log.Fatal("PORT is not set in .env file")
	}

	app := api.NewAPIServer(port)

	if err := app.Run(); err != nil {
		log.Fatal("error to inicialize the server")
	}
}
