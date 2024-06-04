package handlers

import (
	"net/http"

	"github.com/FrancoRutigliano/myMovies/internal/models"
	"github.com/FrancoRutigliano/myMovies/pkg/helpers"
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
	router.HandleFunc("GET /movies", m.GetAllMovies)
}

func (m *MovieHandler) GetAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := m.store.FindAll()
	if err != nil {
		helpers.SendCustom(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.WriteJson(w, http.StatusOK, movies, "movies")
}
