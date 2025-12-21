package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"auto-booking-backend/internal/handlers"
)

func AuthRoutes(r *mux.Router, h *handlers.AuthHandler) {
	r.HandleFunc("/auth/signup", h.Signup).Methods(http.MethodPost)
	r.HandleFunc("/auth/signin", h.Signin).Methods(http.MethodPost)
}