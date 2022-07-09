package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
)

type RepoConfig struct {
	RepoName string `yaml:"name"`
	RepoPath string `yaml:"path"`

	GithubRepoOwner string `yaml:"githubOwner"`
	GithubRepoName  string `yaml:"githubRepoName"`
	GithubRepoHost  string `yaml:"githubRepoHost"`
	BaseBranch      string `yaml:"baseBranch"`
}

type TupperwareConfig struct {
	RepoConfigs []RepoConfig `yaml:"configs"`
}

func EmptyConfig() *TupperwareConfig {
	return &TupperwareConfig{
		RepoConfigs: make([]RepoConfig, 0),
	}
}

func ParseConfig(filepath string) (*TupperwareConfig, error) {
	configFile, err := findConfigFile()
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	return ParseConfigFromBytes(bytes)
}

func ParseConfigFromBytes(bytes []byte) (*TupperwareConfig, error) {
	config := EmptyConfig()

	err := yaml.Unmarshal(bytes, &config)

	if err != nil {
		return nil, err
	}

	return config, nil
}

func findConfigFile() (string, error) {
	dirName, error := os.UserHomeDir()
	if error != nil {
		return "", error
	}
	return filepath.Join(dirName, ".tupperware.yml"), nil
}
