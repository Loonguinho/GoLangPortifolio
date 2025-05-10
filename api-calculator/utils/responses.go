package utils

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

// JSONError logs the error and sends a JSON response with the error message
func JSONError(w http.ResponseWriter, message string, status int, logger *slog.Logger) {
	// Log the error with the provided logger
	logger.Error("Error occurred", slog.String("message", message), slog.Int("status", status))

	// Set the response headers and write the error message in JSON format
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(map[string]string{
		"error": message,
	})
}
