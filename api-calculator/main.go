package main

import (
	"encoding/json"
	"net/http"
)

type CalculatorRequest struct {
	Num1      float64 `json:"num1"`
	Num2      float64 `json:"num2"`
	Operation string  `json:"operation"`
}

type CalculatorResponse struct {
	Result float64 `json:"result"`
	Error  string  `json:"error,omitempty"`
}

func calculate(w http.ResponseWriter, r *http.Request){
	// Set Response Header
	w.Header().Set("Content-Type", "application/json")

	// Allow only POST request
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(CalculatorResponse{Error: "Method not allowed"})
		return
	}
}