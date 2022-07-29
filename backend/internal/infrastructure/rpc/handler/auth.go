package handler

import (
	"context"
	"net/http"

	"github.com/KenethSandoval/doc-md/internal/infrastructure/rpc/pb"
)

type RegisterRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (h *Handler) Register(context.Context, *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	// crear error
	return &pb.RegisterResponse{
		Status: http.StatusAccepted,
		Error:  "error",
	}, nil

	// crear usuario
}
