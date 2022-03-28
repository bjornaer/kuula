package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version number of x tool",
	Long:  `This command can be used get the version number of x tool`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("x v0.0.1-alpha")
	},
}
