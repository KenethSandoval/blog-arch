package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/KenethSandoval/doc-md/internal/config"
	"github.com/KenethSandoval/doc-md/internal/infrastructure/mongodb"
	"github.com/KenethSandoval/doc-md/internal/infrastructure/rpc"
	"github.com/KenethSandoval/doc-md/internal/infrastructure/rpc/pb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("error loading .env file")
	}

	cfg := config.New()

	mongoDB := mongodb.New(cfg)

	if err := mongoDB.Connect(); err != nil {
		log.Fatal("Error connect mongo db")
	}

	listen, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Unable to listen on port :50051: %v", err)
	}

	sHandler := rpc.NewServer(mongoDB)
	sRPC := grpc.NewServer()

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
