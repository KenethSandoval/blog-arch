package port

import (
	"context"

	"github.com/KenethSandoval/doc-md/internal/domain"
)

type AuthUseCase interface {
	Register(context.Context, domain.RegisterPayload) error
	Login(context.Context, domain.Login) (*domain.User, error)
	GetUserByUsername(context.Context, string) (*domain.User, error)
}

type AuthRepository interface {
	Register(context.Context, domain.User) error
	Login(context.Context, domain.Login) (*domain.User, error)
	GetByUsername(context.Context, string) (*domain.User, error)
}
