package cmd

import (
	"github.com/BurntSushi/toml"
)

type server struct {
	Enabled bool
}

type run struct {
	Enabled    bool
	DataSource string
}

type tomlConfig struct {
	Title             string `toml:"title"`
	DependencyHandler string `toml:"dependency_handler"`
	ModelPath         string `toml:"model_path"`
	RequirementsPath  string `toml:"requirements_path"`
	Server            server `toml:"server"`
	Run               run    `toml:"run"`
}

func readTomlConfig(path string) (*tomlConfig, error) {
	var conf tomlConfig
	if _, err := toml.Decode(path, &conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
