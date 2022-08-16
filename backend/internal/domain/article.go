package domain

import (
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrArticleNotFound = errors.New("article not fount")
)

type (
	Article struct {
		ID       primitive.ObjectID `json:"id" bson:"_id"`
		Title    string             `json:"title" bson:"title,omitempty"`
		Content  string             `json:"content" bson:"content,omitempty"`
		NameFile string             `json:"namefile" bson:"namefile,omitempty"`
	}

	ArticleCreatePayload struct {
		ID       primitive.ObjectID `json:"id"`
		Title    string             `json:"title" validate:"required"`
		Content  string             `json:"content" validate:"required"`
		NameFile string             `json:"namefile" validate:"required"`
	}
)

func (Article) TableName() string {
	return "articles"
}
