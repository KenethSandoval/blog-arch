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

	aru := usecase.NewArticleUseCase(arr)

	h := handler.New(handler.HandlerConfig{
		Validator:      v,
		ArticleUseCase: aru,
	})

	e.POST("/articles", h.CreateArticle)

	return e
}
