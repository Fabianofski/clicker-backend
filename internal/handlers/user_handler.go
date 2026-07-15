package handlers

import (
	"net/http"

	"f4b1.dev/clicker-backend/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// @Summary		SignUp
// @Tags			Users
// @Description	Signup with email and password
// @Success		200	{string}	string	"ok"
// @Router			/users/signup [post]
func (h *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	return
}

// @Summary		Login
// @Tags			Users
// @Description	Login with email and password
// @Success		200	{string}	string	"ok"
// @Router			/users/login [post]
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	return
}

// @Summary		Delete Account
// @Description	Delete your account and all your data
// @Tags			Users
// @Success		200	{string}	string	"ok"
// @Success		404	{string}	string	"not found"
// @Router			/users/delete [delete]
func (h *UserHandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	return
}
