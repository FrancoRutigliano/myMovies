package api

import (
	"log"
	"net/http"

	"github.com/FrancoRutigliano/myMovies/pkg/middlewares"
)

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

	middleware := middlewares.MiddlewareChain()

	// SERVER
	log.Println("Listening on port: ", app.addr)

	server := http.Server{
		Addr:    app.addr,
		Handler: middleware(v1),
	}

	return server.ListenAndServe()
}
