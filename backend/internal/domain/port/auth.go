package port

import (
	"context"

	"github.com/KenethSandoval/doc-md/internal/domain"
)

type AuthUseCase interface {
	Register(context.Context, domain.RegisterPayload) error
	GetUserByUsername(context.Context, string) (*domain.Register, error)
}

type AuthRepository interface {
	Register(context.Context, domain.Register) error
	GetByUsername(context.Context, string) (*domain.Register, error)
}
