package handlers

import (
	"fmt"
	"net/http"
)

type MovieHandler struct{}

func (m *MovieHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /movies/{id}", m.GetMovieById)
	router.HandleFunc("POST /movies", m.CreateMovie)
}

func (m *MovieHandler) GetMovieById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "id is: %s", id)
}

func (m *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Movie creada")
}
