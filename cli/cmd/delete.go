package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.AddCommand(deleteWebCmd)
	deleteCmd.AddCommand(deleteAPICmd)
	deleteCmd.AddCommand(deleteDatabaseCmd)

}

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Short:   "delete artifacts (web, api or database)",
	Long:    `This command can be used together with web, api or database sub-commands to delete respective artifacts`,
}

var deleteWebCmd = &cobra.Command{
	Use:   "web",
	Short: "delete web artifacts",
	Long:  `This command can be used to delete web artifacts`,
	Run: func(cmd *cobra.Command, args []string) {
		// *** add code to invoke automation end points below ***
		fmt.Println("Executing 'x delete web' placeholder command")
	},
}

var deleteAPICmd = &cobra.Command{
	Use:   "api",
	Short: "Uneploy API artifacts",
	Long:  `This command can be used to delete API artifacts`,
	Run: func(cmd *cobra.Command, args []string) {
		// *** add code to invoke automation end points below ***
		fmt.Println("Executing 'x delete api' placeholder command")
	},
}

var deleteDatabaseCmd = &cobra.Command{
	Use:   "database",
	Short: "delete database artifacts",
	Long:  `This command can be used to delete database artifacts`,
	Run: func(cmd *cobra.Command, args []string) {
		// *** add code to invoke automation end points below ***
		fmt.Println("Executing 'x delete database' placeholder command")
	},
}
