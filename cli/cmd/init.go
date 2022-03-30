package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(initCmd)

	// initCmd.PersistentFlags().StringP("file", "f", "", "init from file")
	// initCmd.PersistentFlags().StringP("web", "w", "", "init from web url")
	// initCmd.PersistentFlags().StringP("project", "p", "", "init from kuula project")

}

func overwriteExisting(path string) bool {
	// check if exists ~/.kuula/config.yaml config file
	overwrite := true
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return overwrite
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Config file found, do you wish to overwrite it? [Y/n]")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		if text == "n" {
			overwrite = false
		}
	}
	return overwrite
}

func populateConfigFile() error {
	// // prompt user for information
	// // open pop-up to login to our page
	// // get token or something and write to file
	confPath := viper.ConfigFileUsed()
	fmt.Println("Using config file:", confPath)

	if !overwriteExisting(confPath) {
		fmt.Println("keeping current config")
		return nil
	}

	token := "<access-token>"
	viper.Set("access-token", token)
	viper.Set("apiVersion", "v1")
	prompts := []string{"default project:", "username:", "favorite color:"}

	reader := bufio.NewReader(os.Stdin)
	for _, p := range prompts {
		fmt.Print(p)
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("default project:", p) == 0 {
			viper.Set("default-project", text)
			viper.Set("project", text)
		} else if strings.Compare("username:", p) == 0 {
			viper.Set("username", text)
		} else if strings.Compare("favorite color:", p) == 0 {
			fmt.Println("I like that color too!")
		}

	}
	if err := viper.WriteConfig(); err != nil {
		// handle failed write
		return err
	}
	return nil
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init artifacts (web, api or database)",
	Long:  `This command can be used together with web, api or database sub-commands to init respective artifacts`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("let's setup your account")
		err := populateConfigFile()
		handleError(err)
	},
}
