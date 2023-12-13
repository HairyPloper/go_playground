// Package server provides an HTTP server for handling congestion tax calculation requests.
package server

import (
	"congestion-calculator-manager/app/calculator"
	"congestion-calculator-manager/app/helpers"
	taxrules "congestion-calculator-manager/app/tax_rules"
	"congestion-calculator-manager/app/vehicles"
	"errors"

	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// RequestData represents the structure for incoming congestion tax calculation requests.
type RequestData struct {
	Type         string           `json:"type"`
	LicensePlate string           `json:"licenseplate"`
	Dates        []time.Time      `json:"dates"`
	TaxRule      taxrules.TaxRule `json:"tax_rules"`
	IsCustomData bool             `json:"iscustomdata"`
}

// ResultData represents the structure for the result of a congestion tax calculation.
type ResultData struct {
	FeeInfo int
	Error   error
}

var (
	// RequestChannel is a channel for receiving congestion tax calculation requests.
	RequestChannel = make(chan RequestData)
	// ResultChannel is a channel for sending the result of congestion tax calculation.
	ResultChannel = make(chan ResultData)

	// LocalCityCache is a cache for storing city data to avoid reading from disk multiple times.
	LocalCityCache *helpers.Cache = helpers.NewCache()
)

// StartServer initializes and starts the congestion tax calculation server.
func StartServer() {
	http.HandleFunc("/City", customCityTaxHandler)
	http.HandleFunc("/Gothenburg", gothenburgTaxHandler)
	go handleAPiRequests()
	http.ListenAndServe(":8080", nil)
}

// gothenburgTaxHandler handles congestion tax calculation requests for Gothenburg.
func gothenburgTaxHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "GET request received\n")
	case http.MethodPost:
		var requestData RequestData
		err := json.NewDecoder(r.Body).Decode(&requestData)
		if err != nil {
			http.Error(w, "Error decoding JSON", http.StatusBadRequest)
			return
		}
		RequestChannel <- requestData
		requestData.IsCustomData = false
		select {
		case resultInfo := <-ResultChannel:
			if resultInfo.Error != nil {
				http.Error(w, fmt.Sprintf("Error occurred %v", resultInfo.Error), http.StatusInternalServerError)
			} else {
				fmt.Fprintf(w, "Received from server: Type:%v LicensePlate: %v TotalFee:%v\n", requestData.Type, requestData.LicensePlate, resultInfo.FeeInfo)
			}
		case <-time.After(time.Second * 3):
			fmt.Fprintf(w, "request timeout after 3s")
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// customCityTaxHandler handles congestion tax calculation requests for custom cities.
func customCityTaxHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		queryParams := r.URL.Query()
		name := queryParams.Get("name")
		if name == "" {
			http.Error(w, fmt.Sprintln("city name not provided in url"), http.StatusBadRequest)
			break
		}
		//WE can enable cache so that we do not load file from disk each time
		//but then we should set some expiration or way to cleanup/update cache,
		//and remove vehicle information from json

		// jsonSerializedData, exists := LocalCityCache.Get(name)
		// if exists {
		// 	fmt.Println("loaded from cache")
		// } else {
		// 	cityJsonData, err := helpers.ReadContentFromJsonFile(name)
		// 	if err != nil {
		// 		http.Error(w, fmt.Sprintf("Error occurred %v", err), http.StatusInternalServerError)
		// 	}
		// 	jsonSerializedData = cityJsonData
		// 	LocalCityCache.Set(name, jsonSerializedData)
		// }

		cityJsonData, err := helpers.ReadContentFromJsonFile(name)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occurred %v", err), http.StatusInternalServerError)
			break
		}
		cityTaxInfo, err := taxrules.LoadJsonDataForCity(cityJsonData)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error occurred %v", err), http.StatusInternalServerError)
			break
		}

		requestData := RequestData{
			Type:         cityTaxInfo.Vehicle.Type,
			LicensePlate: cityTaxInfo.Vehicle.LicensePlate,
			Dates:        cityTaxInfo.Vehicle.Times,
			TaxRule:      cityTaxInfo.TaxRules,
			IsCustomData: true,
		}
		RequestChannel <- requestData
		select {
		case resultInfo := <-ResultChannel:
			if resultInfo.Error != nil {
				http.Error(w, fmt.Sprintf("Error occurred %v", resultInfo.Error), http.StatusInternalServerError)
				break
			} else {
				fmt.Fprintf(w, "Received from server: Type:%v LicensePlate: %v TotalFee:%v\n", requestData.Type, requestData.LicensePlate, resultInfo.FeeInfo)
			}
		case <-time.After(time.Second * 3):
			fmt.Fprintf(w, "request timeout after 3s")
		}
	case http.MethodPost:
		fmt.Fprintf(w, "POST request received\n")
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// handleAPiRequests continuously listens for incoming congestion tax calculation requests.
func handleAPiRequests() {
	for {
		select {
		case reqData := <-RequestChannel:
			result := ResultData{}
			veh, err := vehicles.GetVehicle(reqData.Type, reqData.LicensePlate)
			if err != nil {
				result.Error = errors.New("error in vehicle information")
			} else {
				result.FeeInfo = calculator.GetTax(
					veh.(vehicles.Vehicle),
					reqData.Dates,
					reqData.IsCustomData,
					reqData.TaxRule)
			}
			ResultChannel <- result
		default:
			fmt.Println("Listening...")
			time.Sleep(time.Second)
		}
	}
}
