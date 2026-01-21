package handlers

import (
    "context"
    "log"

    "github.com/gofiber/fiber/v2"

    "auto-booking-backend/internal/models"
    "auto-booking-backend/internal/services"
)

type AuthHandler struct {
    Service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
    return &AuthHandler{
        Service: service,
    }
}

func (h *AuthHandler) Signup(c *fiber.Ctx) error {
    var req models.SignupRequest

    // Parse request body
    if err := c.BodyParser(&req); err != nil {
        log.Printf("Signup - JSON parse error: %v", err)
        return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
            Error: "Invalid request body",
        })
    }

    // Call service
    user, token, err := h.Service.Signup(context.Background(), req.Name, req.Email, req.Password, req.Role)
    if err != nil {
        log.Printf("Signup failed for email %s: %v", req.Email, err)
        
        // Map service errors to HTTP status codes
        switch err {
        case services.ErrInvalidInput:
            return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
                Error: "All fields are required",
            })
        case services.ErrInvalidEmail:
            return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
                Error: "Invalid email format",
            })
        case services.ErrWeakPassword:
            return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
                Error: "Password must be at least 8 characters",
            })
        case services.ErrInvalidRole:
            return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
                Error: "Invalid role. Must be 'user' or 'admin'",
            })
        case services.ErrEmailExists:
            return c.Status(fiber.StatusConflict).JSON(models.ErrorResponse{
                Error: "Email already exists",
            })
        default:
            return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
                Error: "Internal server error",
            })
        }
    }

    log.Printf("Signup successful for email: %s", req.Email)
    
    // Return success response
    return c.Status(fiber.StatusCreated).JSON(models.AuthResponse{
        Message: "Signup successful",
        Token:   token,
        User:    user,
    })
}

func (h *AuthHandler) Signin(c *fiber.Ctx) error {
    var req models.SigninRequest

    // Parse request body
    if err := c.BodyParser(&req); err != nil {
        log.Printf("Signin - JSON parse error: %v", err)
        return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
            Error: "Invalid request body",
        })
    }

    // Call service
    user, token, err := h.Service.Signin(context.Background(), req.Email, req.Password)
    if err != nil {
        log.Printf("Signin failed for email %s: %v", req.Email, err)
        
        // Map service errors to HTTP status codes
        switch err {
        case services.ErrInvalidInput:
            return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{
                Error: "Email and password are required",
            })
        case services.ErrUserNotFound, services.ErrInvalidPassword:
            return c.Status(fiber.StatusUnauthorized).JSON(models.ErrorResponse{
                Error: "Invalid email or password",
            })
        default:
            return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{
                Error: "Internal server error",
            })
        }
    }

    log.Printf("Signin successful for email: %s", req.Email)
    
    // Return success response
    return c.Status(fiber.StatusOK).JSON(models.AuthResponse{
        Message: "Signin successful",
        Token:   token,
        User:    user,
    })
}