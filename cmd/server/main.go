package main

import (
	"api-contracts-check/cmd/app"
	"api-contracts-check/configs"
	"api-contracts-check/libs/logger"
)

func main() {
	if err := configs.Build(); err != nil {
		logger.Fatal(err)
	}

	logger.NewLogger(configs.Config.LoggerLevel)

	application := app.BuildApplication()
	if err := application.Run(); err != nil {
		panic(err)
	}
}
