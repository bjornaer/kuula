package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)

	// initCmd.PersistentFlags().StringP("file", "f", "", "init from file")
	// initCmd.PersistentFlags().StringP("web", "w", "", "init from web url")
	// initCmd.PersistentFlags().StringP("project", "p", "", "init from kuula project")

}

func createUserConfigDir(path string) (string, error) {
	// find and set .kuula/config.yaml config file
	home, err := homeDir()
	if err != nil {
		return "", err
	}
	wd := filepath.Join(home, ".kuula")
	if path != "" {
		wd = filepath.Join(home, path)
	}
	err = os.Mkdir(wd, 0755)
	if err != nil {
		return "", err
	}
	return wd, nil
}

func createConfigFile() {
	// // prompt user for information
	// // open pop-up to login to our page
	// // get token or something and write to file
	// for {
	// 	fmt.Print("-> ")
	// 	text, _ := reader.ReadString('\n')
	// 	// convert CRLF to LF
	// 	text = strings.Replace(text, "\n", "", -1)

	// 	if strings.Compare("hi", text) == 0 {
	// 		fmt.Println("hello, Yourself")
	// 	}

	// }

}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init artifacts (web, api or database)",
	Long:  `This command can be used together with web, api or database sub-commands to init respective artifacts`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("let's setup your account")
		customPath := args[0]
		confDir, err := createUserConfigDir(customPath)
		handleError(err)

	},
}
