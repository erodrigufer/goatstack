package web

import (
	"fmt"
	"log/slog"
	"net/http"
)

// HandleBadRequest sends a Bad Request response to the client.
func HandleBadRequest(w http.ResponseWriter, errMsg string) {
	http.Error(w, errMsg, http.StatusBadRequest)
}

func HandleOK(w http.ResponseWriter, r *http.Request, errLog *slog.Logger) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte{})
	if err != nil {
		HandleServerError(w, r, fmt.Errorf("unable to send HTTP OK status: %w", err), errLog)
	}
}
