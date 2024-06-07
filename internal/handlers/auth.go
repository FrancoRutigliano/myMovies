package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/FrancoRutigliano/myMovies/internal/models"
	"github.com/FrancoRutigliano/myMovies/pkg/helpers"
	authHelpers "github.com/FrancoRutigliano/myMovies/pkg/helpers/auth"
)

type AuthHandler struct {
	store models.UserAuth
}

func NewAuthHandler(store models.UserAuth) *AuthHandler {
	return &AuthHandler{
		store: store,
	}
}

func (auth *AuthHandler) RegisterRoutes(router *http.ServeMux) {

	router.HandleFunc("POST /auth/register", auth.RegisterUser)
	router.HandleFunc("POST /auth/login", auth.LoginUser)
	router.HandleFunc("PUT /auth/change-password", auth.ChangePassword)
}

func (auth *AuthHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var payload models.UserRegister

	apiErr := helpers.IsValid(r, &payload)
	if apiErr.Msg != "" {
		helpers.SendCustom(w, apiErr.StatusCode, apiErr.Msg)
		return
	}

	// TODO: Si el usuario existe por el email, retornar un error
	if _, ok := auth.store.EmailExist(payload.Email); !ok {
		helpers.SendCustom(w, http.StatusBadRequest, "user already exists")
		return
	}

	// TODO: hash el password
	hashPassword, err := authHelpers.HashPassword(payload.Password)
	if err != nil {
		helpers.SendCustom(w, http.StatusInternalServerError, "oops, something went wrong")
	}

	// TODO: Crear el usuario
	auth.store.CreateUser(&models.User{
		Name:      payload.Name,
		Email:     payload.Email,
		Password:  hashPassword,
		Role:      "user",
		CreatedAt: time.Now().Format(time.RFC3339),
	})
	if err != nil {
		helpers.SendCustom(w, http.StatusInternalServerError, "oops, something went wrong")
	}

	helpers.WriteJson(w, http.StatusCreated, payload, "user")
}

func (auth *AuthHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	// Payload
	var payload models.UserLogin

	apiErr := helpers.IsValid(r, &payload)
	if apiErr.Msg != "" {
		helpers.SendCustom(w, apiErr.StatusCode, apiErr.Msg)
		return
	}

	// TODO: Verificar si el usuario existe
	user, err := auth.store.FindByEmail(payload.Email)
	if err != nil {
		helpers.SendCustom(w, http.StatusBadRequest, err.Error())
		return
	}

	if !authHelpers.ComparePassword(user.Password, []byte(payload.Password)) {
		helpers.SendCustom(w, http.StatusBadRequest, "invalid credentials")
		return
	}

	// JWT token secret
	secret := []byte(os.Getenv("JWT_SECRET"))

	// Generando token
	token, err := authHelpers.CreateJwt(secret, user.Role, user.Email)
	if err != nil {
		helpers.SendCustom(w, http.StatusInternalServerError, "opps, something went wrong")
	}

	helpers.WriteJson(w, http.StatusOK, token, "token")

}

func (auth *AuthHandler) ChangePassword(w http.ResponseWriter, r *http.Request) {
	// Payload
	var payload models.ChangePassword

	apiErr := helpers.IsValid(r, &payload)
	if apiErr.Msg != "" {
		helpers.SendCustom(w, apiErr.StatusCode, apiErr.Msg)
		return
	}

	user, err := auth.store.FindByEmail(payload.Email)
	if err != nil {
		helpers.SendCustom(w, http.StatusBadRequest, err.Error())
		return
	}

	if !authHelpers.ComparePassword(user.Password, []byte(payload.Password)) {
		helpers.SendCustom(w, http.StatusBadRequest, "invalid password")
		return
	}

	hashPassword, err := authHelpers.HashPassword(payload.NewPassword)
	if err != nil {
		helpers.SendCustom(w, http.StatusInternalServerError, "opps, something went wrong")
	}

	err = auth.store.UpdateUserPassword(&models.User{
		ID:       user.ID,
		Password: hashPassword,
	})
	if err != nil {
		helpers.SendCustom(w, http.StatusInternalServerError, "opps, something went wrong")
	}

	helpers.SendCustom(w, http.StatusOK, "password changed successfully")
}
