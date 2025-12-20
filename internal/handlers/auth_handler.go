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

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Name     string
		Email    string
		Password string
		Role     string
	}

	json.NewDecoder(r.Body).Decode(&body)

	err := h.Service.Signup(context.Background(), body.Name, body.Email, body.Password, body.Role)
	if err != nil {
		http.Error(w, "Signup failed", http.StatusBadRequest)
		return
	}

	w.Write([]byte("Signup success"))
}
