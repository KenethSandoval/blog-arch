package handler

import (
	"context"
	"net/http"

	"github.com/KenethSandoval/doc-md/internal/domain"
	"github.com/KenethSandoval/doc-md/internal/infrastructure/rpc/pb"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *Handler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	payload := domain.RegisterPayload{
		ID:       primitive.NewObjectID(),
		Username: req.GetUsername(),
		Password: req.GetPassword(),
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
