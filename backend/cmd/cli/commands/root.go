package commands

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/KenethSandoval/doc-md/internal/infrastructure/rpc/pb"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var rootCmd = &cobra.Command{
	Use:   "cligrpc",
	Short: "A client for grpc server",
	Long:  `A client for grpc server, send a yml data file`,
}

var client pb.AuthServiceClient
var requestCtx context.Context
var requestOpts grpc.DialOption

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	fmt.Println("Starting grpc Cliente")

	requestCtx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	requestOpts = grpc.WithInsecure()

	conn, err := grpc.Dial("localhost:50051", requestOpts)
	if err != nil {
		log.Fatalf("Unable to establish cliente to connection %v", err)
	}

	client = pb.NewAuthServiceClient(conn)
}
