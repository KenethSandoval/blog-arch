package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/KenethSandoval/doc-md/internal/infrastructure/rpc"
	"github.com/KenethSandoval/doc-md/internal/infrastructure/rpc/pb"
	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Unable to listen on port :50051: %v", err)
	}

	sHandler := rpc.NewServer()
	sRPC := grpc.NewServer()

	// pb.RegisterAuthServiceServer(sRPC, sHandler)
	pb.RegisterAuthServiceServer(sRPC, sHandler)

	go func() {
		if err := sRPC.Serve(listen); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()
	fmt.Println("✨ Server succesfully started on port: 50051")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	fmt.Println("💥 shutdown server ...")
	sRPC.Stop()
	listen.Close()
}
