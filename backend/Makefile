# command docker

# development

dev-echo:
	@go run cmd/echo/main.go

dev-cli:
	@go run cmd/cli/main.go

dev-grpc:
	@go run cmd/grpc/main.go

test:
	@go test ./internal/...

proto:
	protoc --go_out=. --go-grpc_out=. api/protobuf/**
