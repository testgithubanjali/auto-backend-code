package models

import "time"

type Driver struct {
	ID          string    `bson:"_id,omitempty" json:"id,omitempty"`
	UserID      string    `bson:"user_id" json:"user_id"`
	IsOnline    bool      `bson:"is_online" json:"is_online"`
	Latitude    float64   `bson:"latitude,omitempty" json:"latitude,omitempty"`
	Longitude   float64   `bson:"longitude,omitempty" json:"longitude,omitempty"`
	LastUpdated time.Time `bson:"last_updated" json:"last_updated"`
}
