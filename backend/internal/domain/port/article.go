package port

import (
	"context"

	"github.com/KenethSandoval/doc-md/internal/domain"
)

type ArticleUseCase interface {
	CreateArticle(context.Context, domain.ArticleCreatePayload) error
}

type ArticleRepository interface {
	Create(context.Context, domain.Article) error
}
