package services

import (
	"context"
	"errors"
	"time"

	"auto-booking-backend/internal/models"
	"auto-booking-backend/internal/repository"
	"auto-booking-backend/internal/utils"
)

type AuthService struct {
	UserRepo  *repository.UserRepo
	JWTSecret string
}

func (as *AuthService) Signup(
	ctx context.Context,
	name string,
	email string,
	password string,
	role string,
) error {

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := models.User{
		Name:      name,
		Email:     email,
		Password:  hashedPassword,
		Role:      role,
		CreatedAt: time.Now(),
	}

	return as.UserRepo.Create(ctx, user)
}

// ✅ THIS FUNCTION WAS MISSING
func (as *AuthService) Signin(
	ctx context.Context,
	email string,
	password string,
) (string, error) {

	user, err := as.UserRepo.FindByEmail(ctx, email)
	if err != nil {
		return "", errors.New("user not found")
	}

	isValid := utils.CheckPasswordHash(user.Password, password)
	if !isValid {
		return "", errors.New("invalid password")
	}

	token, err := utils.GenerateJWT(as.JWTSecret, user.ID, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}