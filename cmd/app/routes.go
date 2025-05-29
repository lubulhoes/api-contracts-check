package app

import (
	"api-contracts-check/cmd/app/utils"
	"api-contracts-check/internals"
	"encoding/json"
	"io"
	"net/http"
)

func (app *Application) registerHandlers() {
	app.mux.HandleFunc("POST /app/create", createAppHandler)
}

func createAppHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		utils.SetResponse(w, http.StatusBadRequest, map[string]string{"error": "Failed to read request body"})
		return
	}

	application := internals.AppDTO{}
	if err := json.Unmarshal(body, &application); err != nil {
		utils.SetResponse(w, http.StatusBadRequest, map[string]string{"error": "Failed to parse request body"})
		return
	}

	if err = internals.CreateAppService(); err != nil {
		utils.SetResponse(w, http.StatusInternalServerError, map[string]string{"error": "Failed to create app"})
		return
	}

	utils.SetResponse(w, http.StatusOK, map[string]string{})
}
