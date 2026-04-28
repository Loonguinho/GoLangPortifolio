package main

import (
	"encoding/json"
	_ "embed"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/loonguinho/api-weather/weather"
)

//go:embed index.html
var htmlContent []byte

type CityRequest struct {
	Name string `json:"name"`
	Lat  string `json:"lat"`
	Long string `json:"long"`
}

type CityResult struct {
	Name    string                      `json:"name"`
	Weather *weather.OpenMeteoResponse `json:"weather"`
	Error   string                      `json:"error,omitempty"`
}

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
	http.HandleFunc("/bulk", getBulkWeatherHandler)
	http.HandleFunc("/history", getHistoryHandler)
	http.HandleFunc("/history/clear", clearHistoryHandler)
	fmt.Println("Starting server on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}

func getBulkWeatherHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Use POST", http.StatusMethodNotAllowed)
		return
	}

	var requestedCities []CityRequest
	if err := json.NewDecoder(r.Body).Decode(&requestedCities); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	var wg sync.WaitGroup
	resultsChan := make(chan CityResult, len(requestedCities))

	startTime := time.Now()
	fmt.Printf("\n🚀 Starting Bulk Fetch for %d cities...\n", len(requestedCities))

	for i, city := range requestedCities {
		wg.Add(1)
		go func(c CityRequest, index int) {
			defer wg.Done()

			workerStart := time.Now()
			fmt.Printf("  🧵 [Worker %d] Starting: %s\n", index+1, c.Name)

			lat, long, name := c.Lat, c.Long, c.Name
			var err error

			if lat == "" || long == "" {
				if name != "" {
					lat, long, name, err = weather.GeocodeCity(name)
					if err != nil {
						resultsChan <- CityResult{Name: c.Name, Error: "City not found"}
						return
					}
				} else {
					resultsChan <- CityResult{Name: "Unknown", Error: "No data"}
					return
				}
			}

			data, err := weather.GetForecast(lat, long)

			result := CityResult{Name: name}
			if err != nil {
				result.Error = err.Error()
			} else {
				result.Weather = data
				_ = weather.SaveWeatherRecord(name, lat, long, data.CurrentWeather.Temperature, data.CurrentWeather.IsDay)
			}

			fmt.Printf("  ✅ [Worker %d] Finished %s in %v\n", index+1, name, time.Since(workerStart))
			resultsChan <- result
		}(city, i)
	}

	go func() {
		wg.Wait()
		close(resultsChan)
		fmt.Printf("🏁 All workers finished. Total time: %v\n\n", time.Since(startTime))
	}()


	var finalResults []CityResult
	for res := range resultsChan {
		finalResults = append(finalResults, res)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(finalResults)
}

func getWeatherHandler(w http.ResponseWriter, r *http.Request) {
	inputName := r.URL.Query().Get("name")
	lat := r.URL.Query().Get("lat")
	long := r.URL.Query().Get("long")

	finalName := inputName

	if lat == "" || long == "" {
		if inputName != "" {
			var err error
			lat, long, finalName, err = weather.GeocodeCity(inputName)
			if err != nil {
				http.Error(w, "City not found", http.StatusNotFound)
				return
			}
		} else {
			http.Error(w, "Name or Coords required", http.StatusBadRequest)
			return
		}
	}

	data, err := weather.GetForecast(lat, long)
	if err != nil {
		http.Error(w, "API Error", http.StatusInternalServerError)
		return
	}

	if finalName == "" { finalName = "Single Search" }
	_ = weather.SaveWeatherRecord(finalName, lat, long, data.CurrentWeather.Temperature, data.CurrentWeather.IsDay)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func getHistoryHandler(w http.ResponseWriter, r *http.Request) {
	records, err := weather.GetHistory()
	if err != nil {
		http.Error(w, "DB Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(records)
}

func clearHistoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "DELETE only", http.StatusMethodNotAllowed)
		return
	}
	weather.ClearHistory()
	w.WriteHeader(http.StatusOK)
}
