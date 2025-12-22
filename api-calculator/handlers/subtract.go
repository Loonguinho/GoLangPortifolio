package handlers

import (
	"log/slog"
	"net/http"
	"strconv"
	"fmt"
)

func SubtractHandler(w http.ResponseWriter, r *http.Request,  logger *slog.Logger) {
	// Log the received parameters
	logger.Info("Received add request", slog.String("method", r.Method))
	if err := r.ParseForm(); err != nil {
		logger.Error("Failed to parse form", slog.String("error", err.Error()))
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	strA := r.FormValue("a")
	strB := r.FormValue("b")
	num1, err1 := strconv.ParseFloat(strA, 64)
	num2, err2 := strconv.ParseFloat(strB, 64)

	if err1 != nil || err2 != nil {
		logger.Error("Invalid number input", slog.String("a", strA), slog.String("b", strB))
		// Return a small HTML error message directly to the user
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("<span style='color: red;'>Invalid input</span>"))
		return
	}
	result := num1 - num2
	logger.Info("Subtraction result", slog.Float64("a", num1), slog.Float64("b", num2), slog.Float64("result", result))

	w.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(w, "Result: <strong>%.2f</strong>", result)
}