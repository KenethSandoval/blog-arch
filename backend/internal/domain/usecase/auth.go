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

// change a authPayload
func (atu *authUseCase) Register(ctx context.Context, payload domain.RegisterPayload) error {
	//TODO: hash password
	data := domain.Register{
		ID:       payload.ID,
		Username: payload.Username,
		Password: payload.Password,
	}

	return atu.authRepo.Register(ctx, data)
}

func (atu *authUseCase) GetUserByUsername(ctx context.Context, username string) (*domain.Register, error) {
	return atu.authRepo.GetByUsername(ctx, username)
}
