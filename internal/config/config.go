package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func Read() (Config, error) {
	configPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, fmt.Errorf("could not find config path: %v", err)
	}

	file, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("error opening json file: %v", err)
	}

	var config Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		return Config{}, fmt.Errorf("error decoding json: %v", err)
	}

	return config, nil
}

func getConfigFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("could not find the home directory of the user: %v", err)
	}
	configPath := filepath.Join(homeDir, configFileName)
	return configPath, nil
}

func (config *Config) SetUser(user string) error {
	if user == "" {
		fmt.Errorf("please add a valid user")
	}
	config.CurrentUserName = user

	configPath, err := getConfigFilePath()
	if err != nil {
		return fmt.Errorf("could not find config path: %v", err)
	}

	file, err := json.Marshal(config)
	if err != nil {
		return fmt.Errorf("error converting config to json: %v", err)
	}

	err = os.WriteFile(configPath, file, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error writing config file to path: %v", err)
	}

	return nil
}
