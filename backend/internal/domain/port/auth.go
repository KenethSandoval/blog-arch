package port

import (
	"context"

	"github.com/KenethSandoval/doc-md/internal/domain"
)

type AuthUseCase interface {
	Register(context.Context, domain.RegisterPayload) error
}

type AuthRepository interface {
	Register(context.Context, domain.Register) error
}
