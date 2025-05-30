package app

import (
	"api-contracts-check/configs"
	"api-contracts-check/libs/logger"
	"fmt"
	"net/http"
)

type Application struct {
	mux        *http.ServeMux
	serverPort string
}

func BuildApplication() *Application {
	mux := http.NewServeMux()
	app := &Application{mux: mux, serverPort: configs.Config.Port}
	app.registerHandlers()

	return app
}

func (app *Application) Run() error {
	logger.Info(fmt.Sprintf("Starting server on port: %s", app.serverPort))
	return http.ListenAndServe(":"+app.serverPort, app.mux)
}
