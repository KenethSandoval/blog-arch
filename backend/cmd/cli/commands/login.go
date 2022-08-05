package commands

import (
	"context"
	"fmt"
	"log"

	"github.com/KenethSandoval/doc-md/cmd/cli/config"
	"github.com/KenethSandoval/doc-md/internal/infrastructure/rpc/pb"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login auth",
	Long:  "",
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := config.ParserConfig()
		if err != nil {
			log.Fatalf("Unable to establish client connection %v", err)
		}

		payload := &pb.LoginRequest{
			Username: data.Body.Username,
			Password: data.Body.Password,
		}

		res, err := client.Login(context.TODO(), payload)

		if err != nil {
			return err
		}
		fmt.Printf("Login successfully %v", res)
		return nil
	},
}

func init() {
	loginCmd.Flags().StringP("file", "f", "", "file for data")
	rootCmd.AddCommand(loginCmd)
}
