package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"auto-booking-backend/internal/config"
	"auto-booking-backend/internal/db"
	"auto-booking-backend/internal/repository"
	"auto-booking-backend/internal/routes"
)

func main() {
	// 1️⃣ Load application config
	cfg := config.LoadConfig()

	// 2️⃣ Connect to MongoDB
	client, err := db.ConnectMongo(cfg.MongoURI)
	if err != nil {
		log.Fatal("❌ Failed to connect to MongoDB:", err)
	}

	// 3️⃣ Get users collection
	usersCollection := client.Database(cfg.DBName).Collection("users")

	// 4️⃣ Initialize repository
	userRepo := repository.NewUserRepo(usersCollection)

	// 5️⃣ Create Fiber app
	app := fiber.New(fiber.Config{
		AppName: "Auto Booking Backend",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			return c.Status(code).JSON(fiber.Map{
				"success": false,
				"error":   err.Error(),
			})
		},
	})

	// 6️⃣ Middleware
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	// 7️⃣ Health check route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "OK",
			"message": "Auto Booking Backend is running 🚀",
		})
	})

	// 8️⃣ Register routes
	routes.RegisterAuthRoutes(app, userRepo, cfg.JWTSecret)

	// 9️⃣ Start server (PORT support for Docker/Cloud)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("🚀 Server running on port", port)
	log.Fatal(app.Listen(":" + port))
}
