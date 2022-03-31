package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var tomlInit = `title: config

dependency-handler: pip
model-path: ./ml_project/model.py
requirements-path: ./ml_project/requirements.txt

server:
  enabled: true

run:
  enabled: false
  data-source: https://kuula.ai/prj-id/data-bucket`

func init() {
	rootCmd.AddCommand(createCmd)

	// initCmd.PersistentFlags().StringP("file", "f", "", "init from file")
	// initCmd.PersistentFlags().StringP("web", "w", "", "init from web url")
	// initCmd.PersistentFlags().StringP("project", "p", "", "init from kuula project")

}

func createProjectConfig(path string) error {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()

	if err != nil {
		return err
	}

	_, err = f.Write([]byte(tomlInit))
	if err != nil {
		return err
	}

	return nil
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "createconfig file for deployment",
	Long:  `This command can be used to create a basic boilerplate config file for a deployment`,
	Run: func(cmd *cobra.Command, args []string) {
		wd, _ := os.Getwd()
		if len(args) != 0 && args[0] != "" {
			wd = filepath.Join(wd, args[0])
		}
		err := os.Chdir(filepath.Join(wd))
		handleError(err)
		err = createProjectConfig(filepath.Join(wd, "config.yaml"))
		handleError(err)
		fmt.Println("Config file created at", wd)

	},
}
