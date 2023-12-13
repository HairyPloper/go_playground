// Package taxrules provides functionality for loading and processing tax rules for congestion tax calculation.

package taxrules

import (
	"congestion-calculator-manager/app/helpers"
	"congestion-calculator-manager/app/vehicles"
	"encoding/json"
	"fmt"
	"time"
)

// TaxRule represents the structure for tax rules used in congestion tax calculation.
type TaxRule struct {
	HourlyPrices       []HourlyPrice `json:"hourly_prices"`
	TaxOnWeekend       bool          `json:"tax_on_weekend"`
	ExcludedMonths     []int         `json:"excluded_months"`
	MaxTaxedFee        int           `json:"max_taxed_fee"`
	ExcludedDates      []time.Time   `json:"excluded_dates"`
	ExcludedDays       []int         `json:"excluded_days"`
	DefaultHourlyPrice int           `json:"default_hourly_price"`
}

// HourlyPrice represents the structure for hourly prices within tax rules.
type HourlyPrice struct {
	StartHour int `json:"start_hour"`
	EndHour   int `json:"end_hour"`
	Rate      int `json:"rate"`
}

// CityData represents the structure for the entire tax rules and vehicle data for a city.
type CityData struct {
	CityName string                  `json:"city_name"`
	Vehicle  vehicles.GenericVehicle `json:"vehicle"`
	TaxRules TaxRule                 `json:"tax_rules"`
}

// LoadJsonDataForCity loads JSON data for a specific city, including tax rules and vehicle information.
// It returns a CityData structure and an error if there's an issue during the process.
func LoadJsonDataForCity(cityName string) (CityData, error) {
	cityData := CityData{}

	// Read JSON content from the file
	jsonData, err := helpers.ReadContentFromJsonFile(cityName)
	if err != nil {
		return cityData, nil
	}

	// Unmarshal JSON data into CityData structure
	err = json.Unmarshal([]byte(jsonData), &cityData)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return cityData, err
	}
	return cityData, nil
}
