package services

import (
	"context"
	"time"

	"auto-booking-backend/internal/models"
	"auto-booking-backend/internal/repository"
	"auto-booking-backend/internal/utils"
)

type AuthService struct {
	UserRepo *repository.UserRepo
}

func (as *AuthService) Signup(ctx context.Context, name, email, password, role string) error {
	hash, _ := utils.HashPassword(password)

	user := models.User{
		Name:      name,
		Email:     email,
		Password:  hash,
		Role:      role,
		CreatedAt: time.Now(),
	}

	return as.UserRepo.Create(ctx, user)
}
