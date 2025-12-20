package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"auto-booking-backend/internal/db"
	"auto-booking-backend/internal/repository"
	"auto-booking-backend/internal/services"
	"auto-booking-backend/internal/handlers"
	"auto-booking-backend/internal/routes"
)

func main() {
	// connect to mongo
	client, err := db.ConnectMongo("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}

	// select database & collection
	collection := client.Database("auto_booking").Collection("users")

	// create layers
	userRepo := &repository.UserRepo{Collection: collection}
	authService := &services.AuthService{UserRepo: userRepo}
	authHandler := &handlers.AuthHandler{Service: authService}

	// setup router
	r := mux.NewRouter()
	routes.RegisterAuthRoutes(r, authHandler)

	log.Println("Server running on port 8080")
	http.ListenAndServe(":8080", r)
}
