package port

import (
	"context"

	"github.com/KenethSandoval/doc-md/internal/domain"
)

type AuthUseCase interface {
	RegisterUser(context.Context, domain.RegisterPayload) error
}

type AuthRepository interface {
	Register(context.Context, domain.User) error
}
