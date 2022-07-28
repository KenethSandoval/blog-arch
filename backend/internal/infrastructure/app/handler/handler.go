package handler

import (
	"github.com/KenethSandoval/doc-md/internal/config"
	"github.com/KenethSandoval/doc-md/internal/domain/port"
)

type (
	Handler struct {
		config *config.Config
		aru    port.ArticleUseCase
	}

	HandlerConfig struct {
		Config         *config.Config
		ArticleUseCase port.ArticleUseCase
	}
)

func New(hc HandlerConfig) *Handler {
	return &Handler{
		config: hc.Config,
		aru:    hc.ArticleUseCase,
	}
}
