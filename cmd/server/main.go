package main

import (
	"api-contracts-check/cmd/app"
	"api-contracts-check/config"
	"log"
	"os"
)

func main() {
	envMode := os.Getenv("ENV_MODE")
	cfg, err := config.LoadConfig(envMode)
	if err != nil {
		log.Fatal(err)
	}

	application := app.BuildApplication(cfg)
	if err := application.Run(); err != nil {
		panic(err)
	}
}
