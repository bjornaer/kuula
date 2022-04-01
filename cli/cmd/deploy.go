package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(deployCmd)

	// deployCmd.AddCommand(deployInferenceCmd)
	deployCmd.PersistentFlags().StringP("file", "f", "", "deploy from file")
	deployCmd.PersistentFlags().StringP("web", "w", "", "deploy from web url")
	deployCmd.PersistentFlags().StringP("project", "p", "", "deploy from kuula project")

}

type DeployPayload struct {
	Username          string
	Project           string
	MLCode            []byte
	Requirements      []byte
	DependencyHandler string
	Server            bool
	Run               bool
	DataSource        string
}

func sendRequest(payload *DeployPayload) {
	jsonPayload, err := json.Marshal(payload)
	req, err := http.NewRequest("POST", "http://localhost:8080/decode", bytes.NewBuffer(jsonPayload))
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
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}

func setPayload(config *yamlConfig) (*DeployPayload, error) {
	var f []byte
	var reqs []byte
	var err error
	fmt.Println(filepath.Join(config.ModelPath))
	f, err = os.ReadFile(filepath.Join(config.ModelPath))
	if err != nil {
		fmt.Println("error reading model file")
		return nil, err
	}
	reqs, err = os.ReadFile(filepath.Join(config.RequirementsPath))
	if err != nil {
		fmt.Println("error reading requirements file")
		return nil, err
	}
	username := viper.GetString("username")
	project := viper.GetString("project")

	return &DeployPayload{
		Username:          username,
		Project:           project,
		MLCode:            f,
		Requirements:      reqs,
		Server:            config.Server.Enabled,
		Run:               config.Run.Enabled,
		DependencyHandler: config.DependencyHandler,
		DataSource:        config.Run.DataSource,
	}, nil

}

var deployCmd = &cobra.Command{
	Use:     "deploy",
	Aliases: []string{"dep", "depl"},
	Short:   "Deploy ML models to your Kuula project",
	Long:    `This command can be used to deploy your ML model to your Kuula Project based on your configuration file`,
	Run: func(cmd *cobra.Command, args []string) {
		// *** add code to invoke automation end points below ***
		var config *yamlConfig
		var err error
		configFileSource, _ := cmd.Flags().GetString("file")
		configWebSource, _ := cmd.Flags().GetString("web")
		projectSource, _ := cmd.Flags().GetString("project")
		// deploymentConfig := getTomlConf()

		if configFileSource != "" {
			config, err = readYamlConfig(configFileSource)
			handleError(err)
		} else if configWebSource != "" {
			// config, err = readConfigFromTheWeb(configWebSource)
			// handleError(err)
			fmt.Println("get source address from the web", configWebSource)

		} else if projectSource != "" {
			// Here there is already a model somewhere in a user owned project, they want to move it
			fmt.Println("get source address from user project", projectSource)
		}
		payload, err := setPayload(config)
		handleError(err)

		// dispatch f to server
		if payload != nil {
			sendRequest(payload)
		}
	},
}

// var deployInferenceCmd = &cobra.Command{
// 	Use:     "inference",
// 	Aliases: []string{"inf", "I"},
// 	Short:   "Deploy inference model artifacts",
// 	Long:    `This command can be used to deploy trained models for inference`,
// }
