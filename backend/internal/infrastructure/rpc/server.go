package rpc

import "github.com/KenethSandoval/doc-md/internal/infrastructure/rpc/handler"

func NewServer() *handler.Handler {
	return handler.New()
}
