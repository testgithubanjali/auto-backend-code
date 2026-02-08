package routes

import (
	"auto-booking-backend/internal/handlers"
	"auto-booking-backend/internal/repository"
	"auto-booking-backend/internal/services"
	"auto-booking-backend/internal/utils"

	"github.com/gofiber/fiber/v2"
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

func RegisterDriverRoutes(app *fiber.App, repo *repository.DriverRepo, jwtSecret string) {
	handler := handlers.NewDriverHandler(repo)

	driver := app.Group("/driver", utils.JWTProtected(jwtSecret))
	driver.Post("/go-online", handler.GoOnline)
	driver.Post("/go-offline", handler.GoOffline)
	driver.Post("/location", handler.UpdateLocation)
}
