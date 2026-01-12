package main

import (
    "fmt"
	"net/http"
	"encoding/json"
	_ "embed"
	"github.com/loonguinho/api-weather/weather"
)

//go:embed index.html
var htmlContent []byte

func main() {
	weather.InitDB()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
            http.NotFound(w, r)
            return
        }
		w.Header().Set("Content-Type", "text/html")
		w.Write(htmlContent)
	})

	http.HandleFunc("/weather", getWeatherHandler)
	http.HandleFunc("/history", getHistoryHandler)
	http.HandleFunc("/history/clear", clearHistoryHandler)
	fmt.Println("Starting server on :8080")
	
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}

func getWeatherHandler(w http.ResponseWriter, r *http.Request) {
	lat := r.URL.Query().Get("lat")
	long := r.URL.Query().Get("long")

	if lat == "" || long == "" {
		http.Error(w, "Latitude and longitude are required", http.StatusBadRequest)
		return
	}

	// Chamamos nosso novo pacote!
	data, err := weather.GetForecast(lat, long)

	if err != nil {
		http.Error(w, "Erro ao buscar clima: "+err.Error(), http.StatusInternalServerError)
		return
	}
	if err := weather.SaveWeatherRecord(lat, long, data.CurrentWeather.Temperature); err != nil {
    	fmt.Println("⚠️ Erro ao salvar histórico:", err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func getHistoryHandler(w http.ResponseWriter, r *http.Request) {
	records, err := weather.GetHistory()
	if err != nil {
		http.Error(w, "Erro ao buscar histórico: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(records)
}

func clearHistoryHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodDelete {
        http.Error(w, "Método não permitido (use DELETE)", http.StatusMethodNotAllowed)
        return
    }

    err := weather.ClearHistory()
    if err != nil {
        http.Error(w, "Erro ao limpar: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Histórico limpo!"))
}