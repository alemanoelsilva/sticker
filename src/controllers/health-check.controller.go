package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type HealthCheckMessage struct {
	Message string
	Uptime  string
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	setHttpHeaders(w)

	fmt.Println("API is running")

	uptimeData, err := os.ReadFile("/proc/uptime")
	if err != nil {
		log.Fatalf("Error reading uptime data: %v", err)
	}
	// Extract the uptime in seconds
	uptimeParts := strings.Fields(string(uptimeData))
	if len(uptimeParts) < 1 {
		log.Fatal("Unexpected format of uptime data")
	}

	uptimeSeconds := uptimeParts[0]
	uptimeSecondsFloat, err := time.ParseDuration(uptimeSeconds + "s")
	if err != nil {
		log.Fatalf("Error parsing uptime data: %v", err)
	}

	// Convert uptime to a human-readable format
	uptime := time.Duration(uptimeSecondsFloat).String()

	healthCheck := HealthCheckMessage{
		Message: "API is running",
		Uptime:  uptime,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(healthCheck)
}
