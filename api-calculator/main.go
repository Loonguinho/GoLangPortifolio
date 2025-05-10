package main

import (
	"api-calculator/handlers"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	// Open or create the log file
	file, err := os.OpenFile("server.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer file.Close()

	// Create the logger with a TextHandler for logging to the file
	handler := slog.NewTextHandler(file, &slog.HandlerOptions{})
	logger := slog.New(handler)
	// Endpoints
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Handling /add request", slog.String("method", r.Method))
		handlers.AddHandler(w, r, logger)
	})
	http.HandleFunc("/subtract", func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Handling /subtract request", slog.String("method", r.Method))
		handlers.SubtractHandler(w, r, logger)
	})
	http.HandleFunc("/multiply", func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Handling /multiply request", slog.String("method", r.Method))
		handlers.MultiplyHandler(w, r, logger)
	})
	http.HandleFunc("/divide", func(w http.ResponseWriter, r *http.Request) {
		logger.Info("Handling /divide request", slog.String("method", r.Method))
		handlers.DivideHandler(w, r, logger)
	})

	fmt.Println("Local server running on :8080")
	logger.Info("Server started", slog.String("address", ":8080"))
	http.ListenAndServe(":8080", nil)
}