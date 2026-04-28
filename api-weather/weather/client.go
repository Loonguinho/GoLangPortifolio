package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type OpenMeteoResponse struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	CurrentWeather CurrentWeather `json:"current_weather"`
}

type GeocodingResponse struct {
	Results []struct {
		Name      string  `json:"name"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Country   string  `json:"country"`
	} `json:"results"`
}

type CurrentWeather struct {
	Temperature float64 `json:"temperature"`
	Windspeed   float64 `json:"windspeed"`
	Winddirection float64 `json:"winddirection"`
	Weathercode int     `json:"weathercode"`
	IsDay       int     `json:"is_day"`
	Time        string  `json:"time"`
}

// GeocodeCity searches for a city name and returns its coordinates and the official name
func GeocodeCity(name string) (string, string, string, error) {
	// Properly escape the name for the URL
	escapedName := url.QueryEscape(name)
	apiURL := fmt.Sprintf("https://geocoding-api.open-meteo.com/v1/search?name=%s&count=1&language=en&format=json", escapedName)
	
	resp, err := http.Get(apiURL)
	if err != nil {
		return "", "", "", err
	}
	defer resp.Body.Close()

	var data GeocodingResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return "", "", "", err
	}

	if len(data.Results) == 0 {
		return "", "", "", fmt.Errorf("city not found")
	}

	result := data.Results[0]
	lat := fmt.Sprintf("%.2f", result.Latitude)
	long := fmt.Sprintf("%.2f", result.Longitude)
	officialName := fmt.Sprintf("%s, %s", result.Name, result.Country)
	
	return lat, long, officialName, nil
}

func GetForecast(lat, long string) (*OpenMeteoResponse, error) {
	if lat == "" || long == "" {
		lat = "-23.55"
		long = "-46.63"
	}

	apiURL := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%s&longitude=%s&current_weather=true", lat, long)

	client := http.Client{Timeout: 60 * time.Second}
	resp, err := client.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("API returned status: %d", resp.StatusCode)
	}

	var data OpenMeteoResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}
