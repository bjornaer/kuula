package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Spins up server to expose an ML model",
	Long:  `This command can be used to spin up a microservice exposing an ML model with a /predict endpoint.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("I make your model be available on host 0.0.0.0:80/predict")
	},
}
