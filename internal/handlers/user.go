package handlers

import (
	"net/http"

	"github.com/FrancoRutigliano/myMovies/internal/models"
	"github.com/FrancoRutigliano/myMovies/pkg/helpers"
	"github.com/FrancoRutigliano/myMovies/pkg/middlewares"
	"github.com/go-playground/validator/v10"
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
	userMiddleware := middlewares.RoleMiddleware("user")
	adminMiddleware := middlewares.RoleMiddleware("admin")
	router.HandleFunc("GET /users", userMiddleware(u.GetAllUsers))
	router.HandleFunc("POST /user/email", adminMiddleware(u.GetUserByEmail))
}

func (u *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := u.store.GetAll()

	helpers.WriteJson(w, http.StatusOK, users, "users")
}

func (u *UserHandler) GetUserByEmail(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email string `json:"email" validate:"required,email"`
	}

	if err := helpers.ParseJson(r, &req); err != nil {
		helpers.SendCustom(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := helpers.Validate.Struct(req); err != nil {
		_ = err.(validator.ValidationErrors)
		helpers.SendCustom(w, http.StatusUnprocessableEntity, err.Error())
		return
	}

	user, err := u.store.FindByEmail(req.Email)
	if err != nil {
		helpers.SendCustom(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.WriteJson(w, http.StatusOK, user, "user")
}
