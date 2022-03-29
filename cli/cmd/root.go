package cmd

import (
	"fmt"
	"os"
	"os/user"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// This is used for config file
	cfgFile string

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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kuula.yaml)")

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

		viper.AddConfigPath(home)
		viper.SetConfigName(".kuula")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
