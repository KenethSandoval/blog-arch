package app

import (
	repository "github.com/KenethSandoval/doc-md/internal/adapter/mongorepo"
	"github.com/KenethSandoval/doc-md/internal/config"
	"github.com/KenethSandoval/doc-md/internal/domain/usecase"
	"github.com/KenethSandoval/doc-md/internal/infrastructure/app/handler"
	"github.com/KenethSandoval/doc-md/internal/infrastructure/mongodb"
	"github.com/KenethSandoval/doc-md/pkg/validation"
	echo "github.com/labstack/echo/v4"
)

func New(cfg *config.Config, dbm *mongodb.MongoDB) *echo.Echo {
	e := echo.New()
	v := validation.New()

	e.Validator = v

	arr := repository.NewArticleMongo(dbm.GetDB())
	art := repository.NewAuthMongo(dbm.GetDB())

	aru := usecase.NewArticleUseCase(arr)
	atu := usecase.NewAuthUseCase(art)

	h := handler.New(handler.HandlerConfig{
		Validator:      v,
		ArticleUseCase: aru,
		AuthUseCase:    atu,
	})

	e.POST("/articles", h.CreateArticle)
	e.POST("/auth/register", h.RegisterUser)

	return e
}
