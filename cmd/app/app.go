package app

import (
	"api-contracts-check/configs"
	"log/slog"
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
	slog.Info("Starting server on port:", app.serverPort)
	return http.ListenAndServe(":"+app.serverPort, app.mux)
}
