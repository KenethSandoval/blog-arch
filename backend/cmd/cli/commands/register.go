package commands

import (
	"context"
	"fmt"
	"log"

	"github.com/KenethSandoval/doc-md/cmd/cli/config"
	"github.com/KenethSandoval/doc-md/internal/infrastructure/rpc/pb"
	"github.com/spf13/cobra"
)

var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "register auth",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := config.ParserConfig()
		if err != nil {
			log.Fatalf("Unable to establish client connection %v", err)
		}

		payload := &pb.RegisterRequest{
			Username: data.Body.Username,
			Password: data.Body.Password,
		}

		res, err := client.Register(context.TODO(), payload)

		if err != nil {
			return err
		}
		fmt.Printf("Register successfully %v", res)
		return nil
	},
}

func init() {
	registerCmd.Flags().StringP("file", "f", "", "file for data")
	rootCmd.AddCommand(registerCmd)
}
