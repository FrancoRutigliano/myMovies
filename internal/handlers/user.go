package handlers

import "net/http"

type UserHandler struct{}

func (u *UserHandler) RegisterRoutes(router *http.ServeMux) {

	router.HandleFunc("POST /users", u.RegisterUser)
	router.HandleFunc("POST /users/login", u.LoginUser)
}

func (u *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
}

func (u *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
}
