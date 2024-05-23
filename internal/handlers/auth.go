package handlers

import "net/http"

type AuthHandler struct{}

func (auth *AuthHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /users/register", auth.RegisterUser)
	router.HandleFunc("POST /users/login", auth.LoginUser)
}

func (auth *AuthHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("register user"))
}

func (auth *AuthHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
}
