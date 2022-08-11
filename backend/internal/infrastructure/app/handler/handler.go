package handler

import (
	"github.com/KenethSandoval/doc-md/internal/config"
	"github.com/KenethSandoval/doc-md/internal/domain/port"
	"github.com/KenethSandoval/doc-md/pkg/validation"
)

type (
	Handler struct {
		config   *config.Config
		validate *validation.Validation
		aru      port.ArticleUseCase
		atu      port.AuthUseCase
	}

	HandlerConfig struct {
		Config         *config.Config
		Validator      *validation.Validation
		ArticleUseCase port.ArticleUseCase
		AuthUseCase    port.AuthUseCase
	}
)

func New(hc HandlerConfig) *Handler {
	return &Handler{
		config:   hc.Config,
		validate: hc.Validator,
		aru:      hc.ArticleUseCase,
		atu:      hc.AuthUseCase,
	}
}
