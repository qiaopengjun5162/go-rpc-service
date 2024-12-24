package routes

import (
	"encoding/json"
	"net/http"
)

const (
	InternalServerError = "Internal server error"
)

// jsonResponse writes the given data as a JSON response to the given http.ResponseWriter.
// It first marshals the given data to JSON, and then writes the resulting JSON data to the
// writer. If any error occurs during this process, it writes an internal server error to the
// writer and returns the error.
//
// Parameters:
//   - w: The http.ResponseWriter to write the JSON response to.
//   - data: The data to be marshalled to JSON and written to the writer.
//   - statusCode: The HTTP status code to be written to the writer.
//
// Returns:
//   - An error, if any error occurs during the process.
func jsonResponse(w http.ResponseWriter, data interface{}, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return err
	}

	w.WriteHeader(statusCode)
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, InternalServerError, http.StatusInternalServerError)
		return err
	}

	return nil
}
