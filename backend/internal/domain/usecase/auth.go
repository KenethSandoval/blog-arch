package usecase

import (
	"context"

	"github.com/KenethSandoval/doc-md/internal/domain"
	"github.com/KenethSandoval/doc-md/internal/domain/port"
)

type authUseCase struct {
	authRepo port.AuthRepository
}

func NewAuthUseCase(atu port.AuthRepository) port.AuthUseCase {
	return &authUseCase{atu}
}

func (atu *authUseCase) Register(ctx context.Context, payload domain.RegisterPayload) error {
	data := domain.User{
		ID:       payload.ID,
		Username: payload.Username,
		Password: payload.Password,
		Role:     payload.Role,
	}

	return atu.authRepo.Register(ctx, data)
}

func (atu *authUseCase) GetUserByUsername(ctx context.Context, username string) (*domain.User, error) {
	return atu.authRepo.GetByUsername(ctx, username)
}
