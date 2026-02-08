package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DriverRepo struct {
	Collection *mongo.Collection
}

func NewDriverRepo(col *mongo.Collection) *DriverRepo {
	return &DriverRepo{Collection: col}
}

func (dr *DriverRepo) GoOnline(ctx context.Context, userID string) error {
	_, err := dr.Collection.UpdateOne(
		ctx,
		bson.M{"user_id": userID},
		bson.M{"$set": bson.M{
			"user_id":      userID,
			"is_online":    true,
			"last_updated": time.Now(),
		}},
		options.Update().SetUpsert(true),
	)
	return err
}

func (dr *DriverRepo) GoOffline(ctx context.Context, userID string) error {
	_, err := dr.Collection.UpdateOne(
		ctx,
		bson.M{"user_id": userID},
		bson.M{"$set": bson.M{
			"is_online":    false,
			"last_updated": time.Now(),
		}},
	)
	return err
}

func (dr *DriverRepo) UpdateLocation(ctx context.Context, userID string, lat, lng float64) error {
	_, err := dr.Collection.UpdateOne(
		ctx,
		bson.M{"user_id": userID, "is_online": true},
		bson.M{"$set": bson.M{
			"latitude":     lat,
			"longitude":    lng,
			"last_updated": time.Now(),
		}},
	)
	return err
}
