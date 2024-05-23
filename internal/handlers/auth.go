package handlers

import (
	"net/http"

	"github.com/FrancoRutigliano/myMovies/internal/models"
	utils "github.com/FrancoRutigliano/myMovies/pkg/helpers"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct{}

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

}

func (auth *AuthHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
}
