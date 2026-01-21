package repository

import (
    "context"

    "auto-booking-backend/internal/models"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

type UserRepo struct {
    Collection *mongo.Collection
}

func NewUserRepo(collection *mongo.Collection) *UserRepo {
    return &UserRepo{
        Collection: collection,
    }
}

func (ur *UserRepo) Create(ctx context.Context, user models.User) error {
    _, err := ur.Collection.InsertOne(ctx, user)
    return err
}

func (ur *UserRepo) FindByEmail(ctx context.Context, email string) (*models.User, error) {
    var user models.User
    err := ur.Collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (ur *UserRepo) EmailExists(ctx context.Context, email string) (bool, error) {
    count, err := ur.Collection.CountDocuments(ctx, bson.M{"email": email})
    if err != nil {
        return false, err
    }
    return count > 0, nil
}