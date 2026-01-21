package routes

import (
    "github.com/gofiber/fiber/v2"
    "auto-booking-backend/internal/handlers"
    "auto-booking-backend/internal/repository"
    "auto-booking-backend/internal/services"
)

func RegisterAuthRoutes(app *fiber.App, userRepo *repository.UserRepo, jwtSecret string) {
    // Initialize service
    authService := services.NewAuthService(userRepo, jwtSecret)
    
    // Initialize handler
    authHandler := handlers.NewAuthHandler(authService)
    
    // Auth routes
    auth := app.Group("/auth")
    auth.Post("/signup", authHandler.Signup)
    auth.Post("/signin", authHandler.Signin)
}