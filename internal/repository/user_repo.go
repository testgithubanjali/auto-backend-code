package repository

import (
	"context"

	"auto-booking-backend/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
	Collection *mongo.Collection
}

func (ur *UserRepo) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := ur.Collection.FindOne(ctx, map[string]string{
		"email": email,
	}).Decode(&user)

	return &user, err
}

func (ur *UserRepo) Create(ctx context.Context, user models.User) error {
	_, err := ur.Collection.InsertOne(ctx, user)
	return err
}