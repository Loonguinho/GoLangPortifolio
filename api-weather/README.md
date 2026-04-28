# Go Weather API - Advanced Concurrency Demo

A professional, high-performance weather dashboard built with Go. This project demonstrates advanced Go features like Goroutines, Channels, and SQLite integration, providing a "Smart Search" experience and concurrent bulk processing.

## 🚀 Advanced Features

*   **Concurrent Bulk Fetch**: Uses **Goroutines and Channels** to fetch weather data for multiple cities simultaneously, reducing total wait time by up to 70%.
*   **Smart Geocoding**: Automatically resolves city names (e.g., "Tokyo") to coordinates using the Open-Meteo Geocoding API.
*   **Real-time Performance Logging**: Backend terminal logs show individual worker execution times and total concurrent processing efficiency.
*   **Modern UI Dashboard**:
    *   **Tabbed Interface**: Switch between single "Smart Search" and "Concurrent Bulk" modes.
    *   **Day/Night Intelligence**: Visual status badges indicating if it's currently day or night in the searched location.
    *   **Floating History Modal**: A persistent, professional modal to view and manage your search history with full context (City Name, Coords, Temp, and Status).
    *   **Interactive Mapping**: Leaflet.js integration for visual location verification.

## 🛠️ Technologies Used

*   **Backend:** Go (Standard Library, `net/http`, `sync`, `encoding/json`)
*   **Concurrency:** Goroutines & Channels
*   **Database:** SQLite3 (via `go-sqlite3`)
*   **Frontend:** Vanilla JS (ES6+), CSS3 Variables, HTML5
*   **APIs:** Open-Meteo (Weather & Geocoding)
*   **Visuals:** Leaflet.js

## 🚦 Getting Started

### Prerequisites

*   Go (version 1.21 or later)
*   CGO (required for `go-sqlite3`)

### Installation & Run

1.  **Clone the repository**:
    ```bash
    git clone https://github.com/loonguinho/api-weather.git
    cd api-weather
    ```
2.  **Run the application**:
    ```bash
    go run main.go
    ```
3.  **Access the Dashboard**:
    Open `http://localhost:8080` in your browser.

## 📡 API Endpoints

| Endpoint | Method | Description |
| :--- | :--- | :--- |
| `/weather` | `GET` | Smart search by `name`, or `lat`/`long` coordinates. |
| `/bulk` | `POST` | Fetches a list of cities concurrently via JSON body. |
| `/history` | `GET` | Retrieves full search history from SQLite. |
| `/history/clear`| `DELETE` | Wipes the search history database. |

### Bulk POST Example
```json
[
  {"name": "São Paulo"},
  {"name": "London", "lat": "51.50", "long": "-0.12"},
  {"name": "Tokyo"}
]
```

## 📈 Learning Outcomes
This project was used to master:
1.  **Concurrency Patterns**: Implementation of `sync.WaitGroup` and buffered channels.
2.  **Closure Scoping**: Safely passing loop variables into goroutines.
3.  **Database Design**: Schema migration and robust data parsing in Go.
4.  **UX/UI Design**: Building responsive, tabbed interfaces with vanilla technology.
