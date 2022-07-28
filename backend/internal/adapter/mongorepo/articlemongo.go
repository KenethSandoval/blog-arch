package repository

import (
	"context"

	"github.com/KenethSandoval/doc-md/internal/domain"
	"github.com/KenethSandoval/doc-md/internal/domain/port"
	"go.mongodb.org/mongo-driver/mongo"
)

type articleMongo struct {
	db *mongo.Database
}

func NewArticleMongo(dbm *mongo.Database) port.ArticleRepository {
	return &articleMongo{dbm}
}

func (adm *articleMongo) Create(ctx context.Context, data domain.Article) error {
	table := data.TableName()

	if _, err := adm.db.Collection(table).InsertOne(ctx, data); err != nil {
		return err
	}

	return nil
}
