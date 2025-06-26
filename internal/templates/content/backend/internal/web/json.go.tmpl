package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// SendJSONResponse transforms a value of any type into JSON and sends the
// JSON data as an HTTP response.
// If the value v cannot be properly coded into JSON the function returns an error.
func SendJSONResponse(w http.ResponseWriter, statusCode int, v any) error {
	// Convert the response value to JSON.
	jsonResponse, err := json.Marshal(v)
	if err != nil {
		return fmt.Errorf("could not marshal value into JSON: %w", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_, err = w.Write(jsonResponse)
	if err != nil {
		return fmt.Errorf("could not write to http.ResponseWriter: %w", err)
	}
	return nil
}
