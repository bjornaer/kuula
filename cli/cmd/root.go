package cmd

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// type Config struct {
// 	ApiVersion     string `mapstructure:"apiVersion"`
// 	Project        string `mapstructure:"project"`
// 	CurrentProject string `mapstructure:"current-project"`
// 	AccessToken    string `mapstructure:"access-token"`
// 	UserName       string `mapstructure:"username"`
// }

var (
	// This is used for config file
	cfgFile string
	// config  Config
	rootCmd = &cobra.Command{
		Use:   "kuula",
		Short: "kuula: A sample CLI application",
		Long:  `kuula: A sample CLI application written in Go as an entry point for the imaginary deployment automation tool`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kuula/config.yaml)")

}

func handleError(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return

}

func homeDir() (string, error) {
	user, err := user.Current()
	if err != nil {
		return "", err
	}
	return user.HomeDir, nil
}

func createUserConfigDir(path string) error {
	// check if exists and set $HOME/.kuula/ dir
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.Mkdir(path, 0755)
		if err != nil {
			return err
		}
	}
	return nil
}

func createUserConfigFile(path string) error {
	// check if exists and set ~/.kuula/config.yaml config file
	if _, err := os.Stat(path); os.IsNotExist(err) {
		f, err := os.Create(path)
		defer f.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

/// this should contain values to connect to our service,
// like token or svc account or whatever
func initConfig() {

	if cfgFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory
		home, err := homeDir()
		handleError(err)
		confHome := filepath.Join(home, ".kuula")
		confName := "config"
		confType := "yaml"
		confPath := filepath.Join(confHome, confName+"."+confType)
		viper.AddConfigPath(confHome)
		viper.SetConfigName(confName)

		err = createUserConfigDir(confHome)
		handleError(err)
		err = createUserConfigFile(confPath)
		handleError(err)

		// if err := viper.SafeWriteConfig(); err != nil {
		// 	// handle failed write
		// 	log.Fatal("Failed writing config to file, please verify file state")
		// }
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		handleError(err)
	}
	// log.Fatal(viper.Get("apiVersion"))
	if viper.Get("apiVersion") == nil {
		fmt.Println("Please run 'kuula init' to set your CLI config")
		// os.Exit(126)
	}
	// viper.Unmarshal(&config)
	// if config.ApiVersion == "" {
	// 	fmt.Println("Please run 'kuula init' to setup your CLI")
	// }
}
