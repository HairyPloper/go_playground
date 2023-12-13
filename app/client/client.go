// Package client provides functionality for sending HTTP requests to the congestion tax calculation server.

package client

import (
	"bytes"
	"congestion-calculator-manager/app/helpers"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// RequestData represents the structure for congestion tax calculation request data.
type RequestData struct {
	Type         string
	LicensePlate string
	Dates        []time.Time
}

// StartClient initializes and sends sample HTTP requests to the congestion tax calculation server.
func StartClient() {

	// Generate 10 random dates for testing
	dates := helpers.GenerateNumberOfDates(10, 2013)
	data := RequestData{
		Type:         "Car",
		Dates:        dates,
		LicensePlate: "ABC123",
	}

	//SendGetRequest("Beograd")
	SendPostRequest(data)
}

// SendGetRequest sends a GET request to the congestion tax calculation server for a specific city.
func SendGetRequest(cityName string) {
	resp, err := http.Get("http://localhost:8080/City?name=" + cityName)
	if err != nil {
		fmt.Println("GET request failed:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Println("GET Response:", string(body))
}

// SendPostRequest sends a POST request to the congestion tax calculation server with the provided request data.
func SendPostRequest(data RequestData) {
	url := "http://localhost:8080/Gothenburg"
	payload, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println("POST request failed:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	fmt.Printf("POST response Status: %v; Body: %v", resp.Status, string(body))
}
