package handler

import (
	"context"
	"net/http"

	"github.com/KenethSandoval/doc-md/internal/domain"
	"github.com/KenethSandoval/doc-md/internal/infrastructure/rpc/pb"
	"github.com/KenethSandoval/doc-md/pkg/bcrypt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {

	user, _ := h.aut.GetUserByUsername(ctx, req.GetUsername())

	if user != nil {
		return &pb.RegisterResponse{
			Status: http.StatusBadRequest,
			Error:  "Username already exists",
		}, nil
	}

	passwordEncrypt, err := bcrypt.HashPassword(req.GetPassword())
	if err != nil {
		return &pb.RegisterResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	payload := domain.RegisterPayload{
		ID:       primitive.NewObjectID(),
		Username: req.GetUsername(),
		Password: passwordEncrypt,
	}

	if err := h.aut.Register(ctx, payload); err != nil {
		return &pb.RegisterResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}
