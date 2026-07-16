package handlers

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"f4b1.dev/clicker-backend/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService, jwt_secret string) *UserHandler {
	return &UserHandler{service: service}
}

// @Summary		SignUp
// @Tags			Users
// @Description	Signup with email and password
// @Accept			json
// @Produce		json
// @Param			request	body		SignUpRequest	true	"sign up information"
// @Success		200		{string}	string			"Success"
// @Failure		400		{string}	string			"Bad Request"
// @Router			/users/signup [post]
func (h *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var req SignUpRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.Email == "" || req.Password == "" || req.Username == "" {
		http.Error(w, "email and password and username are required", http.StatusBadRequest)
		return
	}

	err = h.service.SignUp(r.Context(), req.Email, req.Username, req.Password)
	if err != nil {
		if errors.Is(service.ErrEmailAlreadyInUse, err) {
			http.Error(w, "email already in use", http.StatusBadRequest)
			return
		}
		if errors.Is(service.ErrUsernameAlreadyInUse, err) {
			http.Error(w, "username already taken", http.StatusBadRequest)
			return
		}
		slog.Error("Unexpected Error:", "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Success")
}

// @Summary		Login
// @Tags			Users
// @Description	Login with email and password
// @Accept			json
// @Produce		json
// @Param			request	body		LoginRequest	true	"login information"
// @Success		200		{object}	LoginResponse	"Success"
// @Failure		400		{string}	string			"Bad Request"
// @Failure		401		{string}	string			"Unauthorized"
// @Router			/users/login [post]
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.Email == "" || req.Password == "" {
		http.Error(w, "email and password are required", http.StatusBadRequest)
		return
	}

	token, err := h.service.Login(r.Context(), req.Email, req.Password)
	if err != nil {
		if errors.Is(service.ErrInvalidCredentials, err) {
			http.Error(w, "invalid email or Password.", http.StatusBadRequest)
			return
		}
		slog.Error("Unexpected Error:", "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(LoginResponse{Token: "Bearer " + token})
}

// @Summary		Delete Account
// @Description	Delete your account and all your data
// @Tags			Users
// @Success		200	{string}	string	"ok"
// @Failure		404	{string}	string	"not found"
// @Router			/users/delete [delete]
// @Security		BearerAuth
func (h *UserHandler) DeleteAccount(w http.ResponseWriter, r *http.Request) {
	return
}
