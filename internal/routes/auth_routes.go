package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"auto-booking-backend/internal/handlers"
)

func RegisterAuthRoutes(r *mux.Router, h *handlers.AuthHandler) {
	r.HandleFunc("/signup", h.Signup).Methods(http.MethodPost)
}
