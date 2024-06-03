package handlers

import (
	"net/http"

	"github.com/FrancoRutigliano/myMovies/internal/models"
)

type MovieHandler struct {
	store models.Movies
}

func NewMovieHandler(store models.Movies) *MovieHandler {
	return &MovieHandler{
		store: store,
	}
}

func (m *MovieHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /movies/{id}", m.GetMovieById)
	router.HandleFunc("POST /movies", m.CreateMovie)
}

func (m *MovieHandler) GetMovieById(w http.ResponseWriter, r *http.Request) {

}

func (m *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {

}
