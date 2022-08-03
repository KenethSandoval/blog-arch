package domain

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type (
	User struct {
		ID       primitive.ObjectID `bson:"_id"`
		Username string             `bson:"username"`
		Password string             `bson:"password"`
		Role     string             `bson:"role"`
	}

	RegisterPayload struct {
		ID       primitive.ObjectID `bson:"id"`
		Username string             `bson:"username"`
		Password string             `bson:"password"`
		Role     string             `bson:"role"`
	}
)

func (User) TableName() string {
	return "users"
}
