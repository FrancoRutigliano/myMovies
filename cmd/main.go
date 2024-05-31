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

	app := api.NewAPIServer(config.PORT)

	if err := app.Run(); err != nil {
		log.Fatal("error to inicialize the server")
	}
}
