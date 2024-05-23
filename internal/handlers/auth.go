package handlers

import (
	"net/http"

	"github.com/FrancoRutigliano/myMovies/internal/models"
	"github.com/FrancoRutigliano/myMovies/pkg/helpers"
	utils "github.com/FrancoRutigliano/myMovies/pkg/helpers"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	store models.UserAuth
}

func NewHandler(store models.UserAuth) *AuthHandler {
	return &AuthHandler{
		store: store,
	}
}

func (auth *AuthHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /auth/register", auth.RegisterUser)
	router.HandleFunc("POST /auth/login", auth.LoginUser)
}

func (auth *AuthHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var payload models.UserRegister

	if err := utils.ParseJson(r, &payload); err != nil {
		utils.SendCustom(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := utils.Validate.Struct(&payload); err != nil {
		_ = err.(validator.ValidationErrors)
		utils.SendCustom(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	// TODO: Si el usuario existe por el email, retornar un error
	err := auth.store.EmailExist(payload.Email)
	if err != nil {
		utils.SendCustom(w, http.StatusBadRequest, err.Error())
	}

	// TODO: hash el password
	hashPassword, err := helpers.HashPassword(payload.Password)
	if err != nil {
		helpers.SendCustom(w, http.StatusInternalServerError, err.Error())
	}

}

func (auth *AuthHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
}
