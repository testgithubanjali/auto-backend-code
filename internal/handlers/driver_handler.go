package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"auto-booking-backend/internal/repository"
)

type DriverHandler struct {
	Repo *repository.DriverRepo
}

func NewDriverHandler(repo *repository.DriverRepo) *DriverHandler {
	return &DriverHandler{Repo: repo}
}

func (h *DriverHandler) GoOnline(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	if err := h.Repo.GoOnline(context.Background(), userID); err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(fiber.Map{"message": "Driver is ONLINE"})
}

func (h *DriverHandler) GoOffline(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	if err := h.Repo.GoOffline(context.Background(), userID); err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(fiber.Map{"message": "Driver is OFFLINE"})
}

func (h *DriverHandler) UpdateLocation(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	var body struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}

	if err := c.BodyParser(&body); err != nil {
		return fiber.ErrBadRequest
	}

	err := h.Repo.UpdateLocation(context.Background(), userID, body.Latitude, body.Longitude)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(fiber.Map{"message": "Location updated"})
}
