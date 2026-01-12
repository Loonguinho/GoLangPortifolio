# Go Weather API

This is a simple web application that provides weather information based on latitude and longitude. It uses the Open-Meteo API to fetch weather data and stores the search history in a SQLite database.

## Features

*   Get current weather by latitude and longitude.
*   View a history of previous weather searches.
*   Clear the search history.
*   Interactive map to show the location of the weather search.

## Technologies Used

*   **Backend:** Go
*   **Frontend:** HTML, CSS, JavaScript
*   **Database:** SQLite
*   **API:** Open-Meteo
*   **Map:** Leaflet.js

## Getting Started

### Prerequisites

*   Go (version 1.21 or later)

### Installation

1.  Clone the repository:
    ```bash
    git clone https://github.com/loonguinho/api-weather.git
    ```
2.  Navigate to the project directory:
    ```bash
    cd api-weather
    ```
3.  Install the dependencies:
    ```bash
    go mod tidy
    ```
4.  Run the application:
    ```bash
    go run main.go
    ```
5.  Open your browser and go to `http://localhost:8080`

## API Endpoints

*   `GET /weather?lat=<latitude>&long=<longitude>`: Get the current weather for the specified latitude and longitude.
*   `GET /history`: Get the search history.
*   `DELETE /history/clear`: Clear the search history.

## To Do

*   [ ] Add more weather details (e.g., humidity, pressure).
*   [ ] Add error handling for the frontend.
*   [ ] Dockerize the application.
*   [ ] Add tests.
