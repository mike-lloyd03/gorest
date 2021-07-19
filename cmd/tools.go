package cmd

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Repos []Repo `yaml:"repos"`
}

type Repo struct {
	Name         string `yaml:"name"`
	RepoPath     string `yaml:"repo_path"`
	PasswordFile string `yaml:"password_file"`
}

func ReadConfig(file string) (config *Config, err error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println("Could not open config file for reading. ", err)
		return
	}
	config = &Config{}
	err = yaml.Unmarshal(data, config)
	return
}
