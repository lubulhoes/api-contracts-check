package app

import (
	"net/http"
)

const (
	serverPort = "8080"
)

type Application struct {
	mux        *http.ServeMux
	serverPort string
}

func BuildApplication() *Application {
	mux := http.NewServeMux()
	app := &Application{mux: mux}
	app.registerHandlers()

	return app
}

func (app *Application) Run() error {
	return http.ListenAndServe(":"+serverPort, app.mux)
}
