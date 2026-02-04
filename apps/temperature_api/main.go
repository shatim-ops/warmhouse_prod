package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type TemperatureResponse struct {
	Value       float64   `json:"value"`
	Unit        string    `json:"unit"`
	Timestamp   time.Time `json:"timestamp"`
	Location    string    `json:"location"`
	Status      string    `json:"status"`
	SensorID    string    `json:"sensor_id"`
	SensorType  string    `json:"sensor_type"`
	Description string    `json:"description"`
}

func main() {
	rand.Seed(time.Now().UnixNano())

	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/temperature", temperatureByLocationHandler)
	mux.HandleFunc("/temperature/", temperatureByIDHandler)

	addr := ":8081"
	log.Printf("temperature-api listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func temperatureByLocationHandler(w http.ResponseWriter, r *http.Request) {
	location := r.URL.Query().Get("location")
	if strings.TrimSpace(location) == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "location query param is required"})
		return
	}

	resp := buildResponse(location, "")
	writeJSON(w, http.StatusOK, resp)
}

func temperatureByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/temperature/")
	if strings.TrimSpace(id) == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "sensor id is required"})
		return
	}

	location := "sensor-" + id
	resp := buildResponse(location, id)
	writeJSON(w, http.StatusOK, resp)
}

func buildResponse(location, sensorID string) TemperatureResponse {
	value := randomTemp()
	return TemperatureResponse{
		Value:       value,
		Unit:        "C",
		Timestamp:   time.Now().UTC(),
		Location:    location,
		Status:      statusFromTemp(value),
		SensorID:    sensorID,
		SensorType:  "temperature",
		Description: "Simulated temperature reading",
	}
}

func randomTemp() float64 {
	return -20 + rand.Float64()*60
}

func statusFromTemp(v float64) string {
	switch {
	case v < 0:
		return "cold"
	case v > 30:
		return "hot"
	default:
		return "ok"
	}
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
