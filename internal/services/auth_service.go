package services

import (
	"context"
	"errors"
	"regexp"
	"time"

	"auto-booking-backend/internal/models"
	"auto-booking-backend/internal/repository"
	"auto-booking-backend/internal/utils"
)

type AuthService struct {
	UserRepo  *repository.UserRepo
	JWTSecret string
}

func NewAuthService(userRepo *repository.UserRepo, jwtSecret string) *AuthService {
	return &AuthService{
		UserRepo:  userRepo,
		JWTSecret: jwtSecret,
	}
}

var (
	ErrInvalidInput    = errors.New("invalid input")
	ErrEmailExists     = errors.New("email already exists")
	ErrInvalidEmail    = errors.New("invalid email format")
	ErrWeakPassword    = errors.New("password must be at least 8 characters")
	ErrInvalidRole     = errors.New("invalid role")
	ErrUserNotFound    = errors.New("user not found")
	ErrInvalidPassword = errors.New("invalid password")
)

var validRoles = map[string]bool{
	"user":  true,
	"admin": true,
}

func (as *AuthService) Signup(ctx context.Context, name, email, password, role string) (*models.User, string, error) {
	// Validate inputs
	if name == "" || email == "" || password == "" {
		return nil, "", ErrInvalidInput
	}

	// Validate email format
	if !isValidEmail(email) {
		return nil, "", ErrInvalidEmail
	}

	// Validate password strength
	if len(password) < 8 {
		return nil, "", ErrWeakPassword
	}

	// Validate role (default to "user" if empty or invalid)
	if role == "" {
		role = "user"
	}
	if !validRoles[role] {
		return nil, "", ErrInvalidRole
	}

	// Check if email already exists
	exists, err := as.UserRepo.EmailExists(ctx, email)
	if err != nil {
		return nil, "", err
	}
	if exists {
		return nil, "", ErrEmailExists
	}

	// Hash password
	hash, err := utils.HashPassword(password)
	if err != nil {
		return nil, "", err
	}

	user := models.User{
		Name:      name,
		Email:     email,
		Password:  hash,
		Role:      role,
		CreatedAt: time.Now(),
	}

	err = as.UserRepo.Create(ctx, user)
	if err != nil {
		return nil, "", err
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(as.JWTSecret, user.Email, user.Role)
	if err != nil {
		return nil, "", err
	}

	return &user, token, nil
}

func (as *AuthService) Signin(ctx context.Context, email, password string) (*models.User, string, error) {
	// Validate inputs
	if email == "" || password == "" {
		return nil, "", ErrInvalidInput
	}

	// Find user by email
	user, err := as.UserRepo.FindByEmail(ctx, email)
	if err != nil {
		return nil, "", ErrUserNotFound
	}

	// Check password
	if !utils.CheckPasswordHash(user.Password, password) {
		return nil, "", ErrInvalidPassword
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(as.JWTSecret, user.Email, user.Role)
	if err != nil {
		return nil, "", err
	}

	return user, token, nil
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}
