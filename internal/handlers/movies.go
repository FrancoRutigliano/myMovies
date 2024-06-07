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

func (m *MovieHandler) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("GET /movies", m.GetAllMovies)
	r.HandleFunc("POST /movie", m.CreateMovie)
}

func (m *MovieHandler) GetAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := m.store.FindAll()
	if err != nil {
		helpers.SendCustom(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.WriteJson(w, http.StatusOK, movies, "movies")
}

func (m *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var payload models.Movie

	apiErr := helpers.IsValid(r, &payload)
	if apiErr.Msg != "" {
		helpers.SendCustom(w, apiErr.StatusCode, apiErr.Msg)
		return
	}

	// TODO Verificar si la movie no esta creada ya. Si no esta creada la creamos
	if err := m.store.CreateMovie(&payload); err != nil {
		helpers.SendCustom(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.WriteJson(w, http.StatusCreated, payload, "movie")
}
