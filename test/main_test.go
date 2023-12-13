package test

import (
	"congestion-calculator-manager/app/calculator"
	"congestion-calculator-manager/app/helpers"
	taxrules "congestion-calculator-manager/app/tax_rules"
	"congestion-calculator-manager/app/vehicles"
	"testing"

	"time"
)

func TestCache(t *testing.T) {
	t.Run("TestSetAndGet", testSetAndGet)
}

func testSetAndGet(t *testing.T) {
	cache := helpers.NewCache()

	key := "testKey"
	value := "testValue"

	// Test Set method
	cache.Set(key, value)

	// Test Get method
	result, exists := cache.Get(key)
	if !exists {
		t.Error("Expected value exists in cache, but it doesn't.")
	}

	if result != value {
		t.Errorf("Expected value %s, but got %s", value, result)
	}
}

func TestGenerateRandomDate(t *testing.T) {
	year := 2022
	randomDate := helpers.GenerateRandomDate(year)

	// Ensure the generated date is within the specified year
	if randomDate.Year() != year {
		t.Errorf("Expected year %d, but got %d", year, randomDate.Year())
	}
}

func TestGenerateNumberOfDates(t *testing.T) {
	numDates := 5
	year := 2022
	randomDates := helpers.GenerateNumberOfDates(numDates, year)

	// Ensure the correct number of dates is generated
	if len(randomDates) != numDates {
		t.Errorf("Expected %d dates, but got %d", numDates, len(randomDates))
	}
}

func TestGetTax(t *testing.T) {
	// Create a mock vehicle
	vehicle := vehicles.GenericVehicle{
		LicensePlate: "ABC123",
		Type:         "Car",
		TaxExcluded:  false,
		Times:        []time.Time{time.Now()},
	}

	// Create mock tax rule
	taxRule := taxrules.TaxRule{
		HourlyPrices:       []taxrules.HourlyPrice{{StartHour: 8, EndHour: 18, Rate: 10}},
		TaxOnWeekend:       false,
		ExcludedMonths:     []int{1, 2, 3},
		MaxTaxedFee:        60,
		ExcludedDates:      []time.Time{},
		ExcludedDays:       []int{},
		DefaultHourlyPrice: 7,
	}

	// Create mock date
	dates := []time.Time{time.Now()}

	// Test GetTax method
	result := calculator.GetTax(vehicle, dates, false, taxRule)

	// Validate the result based on your application's logic
	if result < 0 {
		t.Error("Unexpected negative result")
	}
}

func TestIsTollFreeVehicle(t *testing.T) {
	// Mock data
	tollFreeVehicle := vehicles.Car{
		LicensePlate: "FREE123",
	}
	nonTollFreeVehicle := vehicles.Motorbike{
		LicensePlate: "ABC123",
	}

	// Test isTollFreeVehicle function
	tollFreeResult := calculator.IsTollFreeVehicle(tollFreeVehicle)
	nonTollFreeResult := calculator.IsTollFreeVehicle(nonTollFreeVehicle)

	if tollFreeResult {
		t.Error("Expected toll-free vehicle, got non-toll-free")
	}
	if !nonTollFreeResult {
		t.Error("Expected non-toll-free vehicle, got toll-free")
	}
}

func TestIsToolFreeDateWithCustomRules(t *testing.T) {
	// Mock data
	date := time.Date(2023, 1, 1, 12, 0, 0, 0, time.UTC)
	taxRulesWithWeekendExcluded := taxrules.TaxRule{
		TaxOnWeekend: true,
	}
	taxRulesWithMonthExclusion := taxrules.TaxRule{
		ExcludedMonths: []int{1, 2, 3},
	}
	taxRulesWithDayExclusion := taxrules.TaxRule{
		ExcludedDays: []int{1, 15, 30},
	}
	taxRulesWithDateExclusion := taxrules.TaxRule{
		ExcludedDates: []time.Time{
			time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
			time.Date(2023, 2, 14, 0, 0, 0, 0, time.UTC),
		},
	}

	// Test isToolFreeDateWithCustomRules function
	resultWeekend := calculator.IsToolFreeDateWithCustomRules(date, taxRulesWithWeekendExcluded)
	resultMonth := calculator.IsToolFreeDateWithCustomRules(date, taxRulesWithMonthExclusion)
	resultDay := calculator.IsToolFreeDateWithCustomRules(date, taxRulesWithDayExclusion)
	resultDate := calculator.IsToolFreeDateWithCustomRules(date, taxRulesWithDateExclusion)

	if !resultWeekend {
		t.Error("Expected toll-free for weekend, got not toll-free")
	}
	if !resultMonth {
		t.Error("Expected toll-free for excluded month, got not toll-free")
	}
	if !resultDay {
		t.Error("Expected toll-free for excluded day, got not toll-free")
	}
	if !resultDate {
		t.Error("Expected toll-free for excluded date, got not toll-free")
	}
}

func TestIsTollFreeDate(t *testing.T) {
	// Test cases for toll-free dates
	tollFreeDates := []time.Time{
		time.Date(2013, 1, 1, 12, 0, 0, 0, time.UTC),
		time.Date(2013, 3, 28, 12, 0, 0, 0, time.UTC),
		time.Date(2013, 4, 1, 12, 0, 0, 0, time.UTC),
	}

	// Test cases for non-toll-free dates
	nonTollFreeDates := []time.Time{
		time.Date(2013, 1, 2, 12, 0, 0, 0, time.UTC),
		time.Date(2013, 3, 27, 12, 0, 0, 0, time.UTC),
	}

	// Run tests for toll-free dates
	for _, date := range tollFreeDates {
		result := calculator.IsTollFreeDate(date)
		if !result {
			t.Errorf("Expected toll-free for date %v, got not toll-free", date)
		}
	}

	// Run tests for non-toll-free dates
	for _, date := range nonTollFreeDates {
		result := calculator.IsTollFreeDate(date)
		if result {
			t.Errorf("Expected not toll-free for date %v, got toll-free", date)
		}
	}
}
