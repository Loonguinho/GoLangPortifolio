package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func GetForecast(lat, long string) (*OpenMeteoResponse, error) {
	if lat == "" || long == "" {
		lat = "-23.55"
		long = "-46.63"
	}

	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%s&longitude=%s&current_weather=true", lat, long)

	client := http.Client{Timeout: 60 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API retornou status: %d", resp.StatusCode)
	}

	var data OpenMeteoResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}