package repository

import (
	"context"

	"github.com/KenethSandoval/doc-md/internal/domain"
	"github.com/KenethSandoval/doc-md/internal/domain/port"
	"github.com/KenethSandoval/doc-md/pkg/bcrypt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type authMongo struct {
	db *mongo.Database
}

func NewRegisterMongo(dbm *mongo.Database) port.AuthRepository {
	return &authMongo{dbm}
}

func (adm *authMongo) Register(ctx context.Context, data domain.User) error {
	table := data.TableName()

	if _, err := adm.db.Collection(table).InsertOne(ctx, data); err != nil {
		return err
	}
	return nil
}

func (adm *authMongo) Login(ctx context.Context, data domain.Login) (*domain.User, error) {
	result := &domain.User{}
	table := result.TableName()
	filter := bson.M{"username": data.Username}

	err := adm.db.Collection(table).FindOne(ctx, filter).Decode(result)
	if err != nil {
		return nil, err
	}

	if err == mongo.ErrNoDocuments {
		return nil, domain.ErrUserNotFound
	}

	isPassword := bcrypt.CheckPasswordHash(data.Password, result.Password)

	if !isPassword {
		return nil, domain.ErrUserNotFound
	}

	return result, nil
}

// TODO: register for user
func (adm *authMongo) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	result := &domain.User{}
	table := result.TableName()
	filter := bson.M{"username": username}

	err := adm.db.Collection(table).FindOne(ctx, filter).Decode(result)
	if err != nil {
		return nil, err
	}

	if err == mongo.ErrNoDocuments {
		return nil, domain.ErrUserNotFound
	}

	return result, nil
}
