package api

import (
	"log"
	"net/http"

	"github.com/FrancoRutigliano/myMovies/internal/handlers"
	"github.com/FrancoRutigliano/myMovies/internal/service"
	"github.com/FrancoRutigliano/myMovies/pkg/middlewares"
)

// type Handler interface {
// 	RegisterRoutes(router *http.ServeMux)
// }

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (app *APIServer) Run() error {
	// ENRUTADOR
	router := http.NewServeMux()
	v1 := http.NewServeMux()

	v1.Handle("/v1/", http.StripPrefix("/v1/", router))

	moviesHandler := handlers.MovieHandler{}
	moviesHandler.RegisterRoutes(v1)

	// auth
	authStore, err := service.NewUserStore("./data/user.json")
	if err != nil {
		return err
	}
	authHandler := handlers.NewAuthHandler(authStore)
	authHandler.RegisterRoutes(v1)

	// USER
	userStore, err := service.NewUserStore("./data/user.json")
	if err != nil {
		return err
	}
	userHandler := handlers.NewUserHandler(userStore)
	userHandler.RegisterRoutes(v1)

	middleware := middlewares.MiddlewareChain()

	// SERVER
	log.Println("Listening on port: ", app.addr)

	server := http.Server{
		Addr:    app.addr,
		Handler: middleware(v1),
	}

	return server.ListenAndServe()
}
