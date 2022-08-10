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

proto-front:
	protoc api/protobuf/** \
	--js_out=import_style=commonjs,binary:./../client/src/pb \
	--grpc-web_out=import_style=commonjs,mode=grpcwebtext:./../client/src/pb
