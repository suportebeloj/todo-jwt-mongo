package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Profile struct {
	FirstName string `json:"first_name" bson:"first_name"`
	LastName  string `json:"last_name" bson:"last_name"`
	Email     string `json:"email" bson:"email"`
	ImageUrl  string `json:"image_url" bson:"image_url"`
}

type User struct {
	ID             primitive.ObjectID `json:"id" bson:"_id"`
	Username       string             `json:"username" bson:"username"`
	HashedPassword string             `json:"hashed_password" bson:"hashed_password"`
	Salt           int                `json:"salt" bson:"salt"`
	Profile        Profile            `json:"profile" bson:"profile"`
	CreatedAt      time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at" bson:"updated_at"`
	DeletedAt      time.Time          `json:"deleted_at" bson:"deleted_at"`
}

type AuthUser struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type RegisterUser struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Email    string `json:"email" bson:"email"`
}
