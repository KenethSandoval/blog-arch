package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	Register struct {
		ID       primitive.ObjectID `bson:"_id"`
		Username string             `bson:"username"`
		Password string             `bson:"password"`
	}

	RegisterPayload struct {
		ID       primitive.ObjectID `bson:"id"`
		Username string             `bson:"username"`
		Password string             `bson:"password"`
	}
)

func (Register) TableName() string {
	return "users"
}
