package handler

import (
	"context"
	"net/http"

	"github.com/KenethSandoval/doc-md/internal/domain"
	"github.com/KenethSandoval/doc-md/internal/infrastructure/middle"
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
		Role:     req.GetRole(),
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

func (h *Handler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	payload := domain.Login{
		Username: req.Username,
		Password: req.Password,
	}

	user, err := h.aut.Login(ctx, payload)

	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusUnauthorized,
			Error:  err.Error(),
		}, nil
	}

	token, err := middle.CreateJWTToken("secret", user)

	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}
