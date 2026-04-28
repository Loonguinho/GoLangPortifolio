package weather

import (
	"database/sql"
	"log"
	"time"
	"fmt"
	"strconv"
	"strings"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type WeatherRecord struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	Temperature float64   `json:"temperature"`
	IsDay       int       `json:"is_day"`
	Timestamp   time.Time `json:"timestamp"`
}

func InitDB() {
	var err error

	db, err = sql.Open("sqlite3", "weather.db")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	// Updated table with name and is_day
	createTableSQL := `CREATE TABLE IF NOT EXISTS weather_records (
        "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "name" TEXT,
        "latitude" REAL,
        "longitude" REAL,
        "temperature" REAL,
        "is_day" INTEGER,
        "timestamp" DATETIME
    );`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
}

func SaveWeatherRecord(name string, lat string, long string, temp float64, isDay int) error {
	lat = strings.TrimSpace(strings.Replace(lat, ",", ".", -1))
	long = strings.TrimSpace(strings.Replace(long, ",", ".", -1))

	latFloat, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		return fmt.Errorf("error parsing latitude '%s': %v", lat, err)
	}

	longFloat, err := strconv.ParseFloat(long, 64)
	if err != nil {
		return fmt.Errorf("error parsing longitude '%s': %v", long, err)
	}

    query := "INSERT INTO weather_records(name, latitude, longitude, temperature, is_day, timestamp) VALUES(?, ?, ?, ?, ?, ?)"
    _, err = db.Exec(query, name, latFloat, longFloat, temp, isDay, time.Now())
    if err != nil {
        return fmt.Errorf("error executing insert: %v", err)
    }

    return nil
}

func GetHistory() ([]WeatherRecord, error) {
	rows, err := db.Query("SELECT id, name, latitude, longitude, temperature, is_day, timestamp FROM weather_records ORDER BY timestamp DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []WeatherRecord
	for rows.Next() {
		var record WeatherRecord
		err := rows.Scan(&record.ID, &record.Name, &record.Latitude, &record.Longitude, &record.Temperature, &record.IsDay, &record.Timestamp)
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	return records, nil
}

func ClearHistory() error {
	_, err := db.Exec("DELETE FROM weather_records")
	return err
}
