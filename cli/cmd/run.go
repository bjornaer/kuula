package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(runCmd)
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a prediction once",
	Long:  `This command can be used to run a prediciton for a given model over a given data set once`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("running model prediction now...")
	},
}
