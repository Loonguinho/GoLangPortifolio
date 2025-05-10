package handlers

import (
	"api-calculator/models"
	"api-calculator/utils"
	"encoding/json"
	"log/slog"
	"net/http"
)

func DivideHandler(w http.ResponseWriter, r *http.Request,  logger *slog.Logger) {
	// Log the received parameters
	logger.Info("Received divide request", slog.String("method", r.Method))
	var nums models.Numbers
	err := json.NewDecoder(r.Body).Decode(&nums)

	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	if nums.Num2 == 0 {
		utils.JSONError(w, "Cannot divide by zero", http.StatusBadRequest, logger)
		return
	}

	result := nums.Num1 / nums.Num2
	json.NewEncoder(w).Encode(map[string]float64{"result": result})

}