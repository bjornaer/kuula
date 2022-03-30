package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)

	// initCmd.PersistentFlags().StringP("file", "f", "", "init from file")
	// initCmd.PersistentFlags().StringP("web", "w", "", "init from web url")
	// initCmd.PersistentFlags().StringP("project", "p", "", "init from kuula project")

}

func createUserConfig(path string) {
	// find and set .kuula.yaml config file
}

var initCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"dep", "depl"},
	Short:   "init artifacts (web, api or database)",
	Long:    `This command can be used together with web, api or database sub-commands to init respective artifacts`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("let's setup your account")
	},
}
