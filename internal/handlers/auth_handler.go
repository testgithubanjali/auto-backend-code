package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"auto-booking-backend/internal/services"
)

type AuthHandler struct {
	Service *services.AuthService
}

func (ah *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	err := ah.Service.Signup(
		context.Background(),
		req.Name,
		req.Email,
		req.Password,
		req.Role,
	)

	if err != nil {
		http.Error(w, "Signup failed", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Signup successful"))
}

func (ah *AuthHandler) Signin(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	json.NewDecoder(r.Body).Decode(&req)

	token, err := ah.Service.Signin(
		context.Background(),
		req.Email,
		req.Password,
	)

	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	w.Write([]byte(token))
}