package handler

import (
	"github.com/KenethSandoval/doc-md/internal/domain/port"
	"github.com/KenethSandoval/doc-md/internal/infrastructure/rpc/pb"
)

type (
	Handler struct {
		pb.AuthServiceServer
		aut port.AuthUseCase
	}

	HandlerConfig struct {
		AuthUseCase port.AuthUseCase
	}
)

func New(hc HandlerConfig) *Handler {
	return &Handler{
		aut: hc.AuthUseCase,
	}
}
