package main

import (
    "fmt"
	"net/http"
	"encoding/json"
	"time"
)

type OpenMeteoResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	CurrentWeather CurrentWeather `json:"current_weather"`
}

type CurrentWeather struct {
	Temperature float64 `json:"temperature"`
	Windspeed   float64 `json:"windspeed"`
	Winddirection float64 `json:"winddirection"`
	Weathercode int     `json:"weathercode"`
	IsDay       int     `json:"is_day"`
	Time        string  `json:"time"`
}

func main() {
	http.HandleFunc("/weather", getWeatherHandler)
	fmt.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}

func getWeatherHandler(w http.ResponseWriter, r *http.Request) {
	lat := "-23.55"
	long := "-46.63"
	//Build api link
	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%s&longitude=%s&current_weather=true", lat, long)
	fmt.Println("Fetching weather data from:", url)
	client := &http.Client{Timeout: 60 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		http.Error(w, "Failed to fetch weather data"+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var weatherResponse OpenMeteoResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResponse); err != nil {
		http.Error(w, "Failed to parse weather data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(weatherResponse)}