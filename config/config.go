package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Port        string `json:"port"`
	LoggerLevel string `json:"logger_level"`
}

func LoadConfig(envMode string) (*Config, error) {
	if envMode != "development" {
		return nil, nil
	}

	file, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
