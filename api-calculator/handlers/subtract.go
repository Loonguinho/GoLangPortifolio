package handlers

import (
	"api-calculator/models"
	"encoding/json"
	"log/slog"
	"net/http"
)

func SubtractHandler(w http.ResponseWriter, r *http.Request,  logger *slog.Logger) {
	// Log the received parameters
	logger.Info("Received subtract request", slog.String("method", r.Method))

	var nums models.Numbers
	err := json.NewDecoder(r.Body).Decode(&nums)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	result := nums.Num1 - nums.Num2
	json.NewEncoder(w).Encode(map[string]float64{"result": result})
}