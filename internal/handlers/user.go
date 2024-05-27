package handlers

import (
	"net/http"

	"github.com/FrancoRutigliano/myMovies/internal/models"
	"github.com/FrancoRutigliano/myMovies/pkg/helpers"
)

type UserHandler struct {
	store models.Users
}

func NewUserHandler(store models.Users) *UserHandler {
	return &UserHandler{
		store: store,
	}
}

func (u *UserHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /users", u.GetAllUsers)
}

func (u *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := u.store.GetAll()

	helpers.WriteJson(w, http.StatusOK, users, "users")
}
