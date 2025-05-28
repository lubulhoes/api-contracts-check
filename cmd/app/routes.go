package app

import (
	"api-contracts-check/cmd/app/utils"
	"net/http"
)

func (app *Application) registerHandlers() {
	app.mux.HandleFunc("POST /api/new", app.newApiHandler)
}

func (app *Application) newApiHandler(w http.ResponseWriter, r *http.Request) {
	utils.SetResponse(w, http.StatusOK, map[string]string{})
}
