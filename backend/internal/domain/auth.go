package domain

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrUserNotFound      = errors.New("User not found")
	ErrUserAlreadyExists = errors.New("User already exists")
)

type (
	User struct {
		ID       primitive.ObjectID `json:"id" bson:"_id"`
		Username string             `json:"username" bson:"username"`
		Password string             `json:"password" bson:"password"`
		Role     string             `json:"role" bson:"role"`
	}

	RegisterPayload struct {
		ID       primitive.ObjectID `json:"id" bson:"_id"`
		Username string             `json:"username" bson:"username"`
		Password string             `json:"password" bson:"password"`
		Role     string             `json:"role" bson:"role"`
	}

	LoginPayload struct {
		Username string `json:"username" bson:"username"`
		Password string `json:"password" bson:"password"`
	}

	LoginResponse struct {
		User  User   `json:"user"`
		Token string `json:"token"`
	}
)

func (User) TableName() string {
	return "users"
}
