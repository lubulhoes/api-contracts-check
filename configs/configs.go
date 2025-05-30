package configs

import (
	"encoding/json"
	"io"
	"os"
)

const (
	development  = "development"
	relativePath = "resources/configuration.json"
)

var Config *AppConfig

type AppConfig struct {
	Port        string `json:"port"`
	LoggerLevel string `json:"logger_level"`
}

func Build() error {
	env := os.Getenv("GO_ENVIRONMENT")
	if env == development {
		return ReadConfigJSON(relativePath)
	}
	file, err := os.Open("configs.json")
	if err != nil {
		return err
	}
	defer file.Close()

	if err := json.NewDecoder(file).Decode(&Config); err != nil {
		return err
	}

	return nil
}

func ReadConfigJSON(path string) error {
	jsonFile, err := os.Open(path)
	if err != nil {
		return err
	}

	byteArray, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	return json.Unmarshal(byteArray, &Config)
}
