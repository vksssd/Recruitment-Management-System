package schema

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Register struct {
	Id primitive.ObjectID `json:"_id" bson:"_id"`
	UserID string `json:"user_id" bson:"user_id"`
	Name string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	PasswordHash string `json:"password_hash" bson:"password_hash"`
	UserType string `json:"user_type" bson:"user_type"`
	ProfileHeading string `json:"profile_heading" bson:"profile_heading"`
	Address string `json:"address" bson:"address"`
	AccessToken string `json:"access_token" bson:"access_token"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdateAt time.Time `json:"updated_at" bson:"updated_at"`
}

type User struct {
	Name string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
	PasswordHash string `json:"password_hash" bson:"password_hash"`
	UserType string `json:"user_type" bson:"user_type"`
	ProfileHeading string `json:"profile_heading" bson:"profile_heading"`
	Address string `json:"address" bson:"address"`
}
