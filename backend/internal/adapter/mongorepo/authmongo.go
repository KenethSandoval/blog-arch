package repository

import (
	"context"

	"github.com/KenethSandoval/doc-md/internal/domain"
	"github.com/KenethSandoval/doc-md/internal/domain/port"
	"go.mongodb.org/mongo-driver/mongo"
)

type authMongo struct {
	db *mongo.Database
}

func NewAuthMongo(dbm *mongo.Database) port.AuthRepository {
	return &authMongo{dbm}
}

func (adm *authMongo) Register(ctx context.Context, data domain.User) error {
	table := data.TableName()

	if _, err := adm.db.Collection(table).InsertOne(ctx, data); err != nil {
		return err
	}

	return nil
}
