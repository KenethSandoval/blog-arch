package usecase

import (
	"context"

	"github.com/KenethSandoval/doc-md/internal/domain"
	"github.com/KenethSandoval/doc-md/internal/domain/port"
)

type articleUseCase struct {
	articleRepo port.ArticleRepository
}

func NewArticleUseCase(aru port.ArticleRepository) port.ArticleUseCase {
	return &articleUseCase{aru}
}

func (aru *articleUseCase) CreateArticle(ctx context.Context, payload domain.ArticleCreatePayload) error {
	data := domain.Article{
		ID:      payload.ID,
		Title:   payload.Title,
		Content: payload.Content,
	}
	return aru.articleRepo.Create(ctx, data)
}
