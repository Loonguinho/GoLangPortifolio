package weather

import (
	"database/sql"
	"log"
	"time"
	"fmt"
	"strconv"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

type WeatherRecord struct {
	ID          int       `json:"id"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	Temperature float64   `json:"temperature"`
	Timestamp   time.Time `json:"timestamp"`
}

func InitDB() {
	var err error

	db, err = sql.Open("sqlite3", "weather.db")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS weather_records (
        "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        "latitude" REAL,
        "longitude" REAL,
        "temperature" REAL,
        "timestamp" DATETIME
    );`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal("Failed to create table:", err)
	}
}

func SaveWeatherRecord(lat string, long string, temp float64) error {
    fmt.Println("💾 TENTANDO SALVAR NO BANCO...") // Debug 1

	latFloat, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		fmt.Println("🚨 ERRO AO CONVERTER LATITUDE:", err)
		return err
	}

	longFloat, err := strconv.ParseFloat(long, 64)
	if err != nil {
		fmt.Println("🚨 ERRO AO CONVERTER LONGITUDE:", err)
		return err
	}

    query := "INSERT INTO weather_records(latitude, longitude, temperature, timestamp) VALUES(?, ?, ?, ?)"
    stmt, err := db.Prepare(query)
    if err != nil {
        fmt.Println("🚨 ERRO AO PREPARAR QUERY:", err) // Debug 2
        return err
    }
    defer stmt.Close()

    _, err = stmt.Exec(latFloat, longFloat, temp, time.Now())
    if err != nil {
        fmt.Println("🚨 ERRO AO EXECUTAR INSERT:", err) // Debug 3
        return err
    }

    fmt.Println("✅ SUCESSO! DADOS GRAVADOS NO ARQUIVO.") // Debug 4
    return nil
}

func GetHistory() ([]WeatherRecord, error) {
	rows, err := db.Query("SELECT id, latitude, longitude, temperature, timestamp FROM weather_records ORDER BY timestamp DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var records []WeatherRecord
	for rows.Next() {
		var record WeatherRecord
		err := rows.Scan(&record.ID, &record.Latitude, &record.Longitude, &record.Temperature, &record.Timestamp)
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