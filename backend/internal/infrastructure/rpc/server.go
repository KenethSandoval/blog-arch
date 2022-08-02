package rpc

import (
	repository "github.com/KenethSandoval/doc-md/internal/adapter/mongorepo"
	"github.com/KenethSandoval/doc-md/internal/domain/usecase"
	"github.com/KenethSandoval/doc-md/internal/infrastructure/mongodb"
	"github.com/KenethSandoval/doc-md/internal/infrastructure/rpc/handler"
)

func NewServer(dbm *mongodb.MongoDB) *handler.Handler {

	arr := repository.NewRegisterMongo(dbm.GetDB())

	atu := usecase.NewAuthUseCase(arr)

	return handler.New(handler.HandlerConfig{
		AuthUseCase: atu,
	})
}
