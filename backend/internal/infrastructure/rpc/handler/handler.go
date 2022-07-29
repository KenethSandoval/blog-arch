package handler

import "github.com/KenethSandoval/doc-md/internal/infrastructure/rpc/pb"

type Handler struct {
	pb.AuthServiceServer
}

func New() *Handler {
	return &Handler{}
}
