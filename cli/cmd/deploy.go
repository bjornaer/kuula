package cmd

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deployCmd)

	deployCmd.AddCommand(deployInferenceCmd)
	deployCmd.AddCommand(deployAPICmd)
	deployCmd.AddCommand(deployDatabaseCmd)
	deployCmd.PersistentFlags().StringP("file", "f", "", "deploy from file")
	deployCmd.PersistentFlags().StringP("web", "w", "", "deploy from web url")
	deployCmd.PersistentFlags().StringP("project", "p", "", "deploy from x project")

}

func sendRequest(bod []byte) {
	req, err := http.NewRequest("POST", "http://localhost:8090/decode", bytes.NewBuffer(bod))
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
}

var deployCmd = &cobra.Command{
	Use:     "deploy",
	Aliases: []string{"dep", "depl"},
	Short:   "Deploy artifacts (web, api or database)",
	Long:    `This command can be used together with web, api or database sub-commands to deploy respective artifacts`,
}

var deployInferenceCmd = &cobra.Command{
	Use:     "inference",
	Aliases: []string{"inf", "I"},
	Short:   "Deploy inference model artifacts",
	Long:    `This command can be used to deploy trained models for inference`,
	Run: func(cmd *cobra.Command, args []string) {
		// *** add code to invoke automation end points below ***
		var f []byte
		var err error
		fileSource, _ := cmd.Flags().GetString("file")
		webSource, _ := cmd.Flags().GetString("web")
		projectSource, _ := cmd.Flags().GetString("project")

		if fileSource != "" {
			f, err = os.ReadFile(fileSource)
			handleError(err)
		} else if webSource != "" {
			fmt.Println("get source address from the web", webSource)

		} else if projectSource != "" {
			fmt.Println("get source address from user project", projectSource)
		}
		fmt.Println("Executing 'x deploy inference' placeholder command")

		// dispatch f to server
		if f != nil {
			sendRequest(f)
		}
	},
}

var deployAPICmd = &cobra.Command{
	Use:   "api",
	Short: "Deploy API artifacts",
	Long:  `This command can be used to deploy API artifacts`,
	Run: func(cmd *cobra.Command, args []string) {
		// *** add code to invoke automation end points below ***
		fmt.Println("Executing 'x' deploy api' placeholder command")
	},
}

var deployDatabaseCmd = &cobra.Command{
	Use:   "database",
	Short: "Deploy database artifacts",
	Long:  `This command can be used to deploy database artifacts`,
	Run: func(cmd *cobra.Command, args []string) {
		// *** add code to invoke automation end points below ***
		fmt.Println("Executing 'x deploy database' placeholder command")
	},
}
