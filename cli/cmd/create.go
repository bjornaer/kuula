package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var tomlInit = `title = "config"
[server]
enable = true
model_path = "./model.py"
requirements_path = "./requirements.txt"

[run]
enable = false`

func init() {
	rootCmd.AddCommand(createCmd)

	// initCmd.PersistentFlags().StringP("file", "f", "", "init from file")
	// initCmd.PersistentFlags().StringP("web", "w", "", "init from web url")
	// initCmd.PersistentFlags().StringP("project", "p", "", "init from kuula project")

}

func createProjectConfig(path string) {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(path+".toml", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	_, err = f.Write([]byte(tomlInit))
	if err != nil {
		log.Fatal(err)
	}

	f.Close()
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "createconfig file for deployment",
	Long:  `This command can be used to create a basic boilerplate config file for a deployment`,
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		if args[0] != "" {
			wd = filepath.Join(wd, args[0])
		}
		err := os.Chdir(filepath.Join(wd))
		handleError(err)
	},
}