package app

import (
	"api-contracts-check/config"
	"net/http"
)

type Application struct {
	mux        *http.ServeMux
	serverPort string
}

func BuildApplication(cfg *config.Config) *Application {
	mux := http.NewServeMux()
	app := &Application{mux: mux, serverPort: cfg.Port}
	app.registerHandlers()

	return app
}

func (app *Application) Run() error {
	return http.ListenAndServe(":"+app.serverPort, app.mux)
}
