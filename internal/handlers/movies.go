package handlers

import (
	"net/http"
	"time"

	"github.com/FrancoRutigliano/myMovies/internal/models"
	utils "github.com/FrancoRutigliano/myMovies/pkg/helpers"
	"github.com/go-playground/validator/v10"
)

type MovieHandler struct{}

func (m *MovieHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /movies/{id}", m.GetMovieById)
	router.HandleFunc("POST /movies", m.CreateMovie)
}

func (m *MovieHandler) GetMovieById(w http.ResponseWriter, r *http.Request) {
	id, err := utils.ReadIdParam(r)
	if err != nil {
		utils.SendCustom(w, http.StatusBadRequest, err.Error())
		return
	}

	// movie model

	data := models.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casa blanca",
		Runtime:   102,
		Genres:    []string{"drama", "comedia", "horror"},
		Version:   1,
	}

	if err = utils.WriteJson(w, http.StatusOK, data, "movie"); err != nil {
		utils.SendCustom(w, http.StatusBadRequest, err.Error())
		return
	}

}

func (m *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var payload models.Movie

	if err := utils.ParseJson(r, &payload); err != nil {
		utils.SendCustom(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		_ = err.(validator.ValidationErrors)
		utils.SendCustom(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	utils.WriteJson(w, http.StatusCreated, payload, "movie")
}
