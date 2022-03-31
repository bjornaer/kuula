package cmd

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

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
	var c yamlConfig
	// _, b, _, _ := runtime.Caller(0)
	// basepath := filepath.Dir(b)
	// place := filepath.Join(basepath, path)
	// fmt.Println(place)
	yamlFile, err := ioutil.ReadFile(filepath.Join(path))
	if err != nil {
		fmt.Println("read error:")
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		fmt.Println("Unmarshal error:")
		return nil, err
	}

	return &c, nil
}
