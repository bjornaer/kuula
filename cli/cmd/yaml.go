package cmd

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Server struct {
	Enabled bool
}

type Run struct {
	Enabled    bool
	DataSource string
}

type yamlConfig struct {
	Title             string
	DependencyHandler string
	ModelPath         string
	RequirementsPath  string
	Server            Server `yaml:"server"`
	Run               Run    `yaml:"run"`
}

func readYamlConfig(path string) (*yamlConfig, error) {
	var c *yamlConfig
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
