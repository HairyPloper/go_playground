// Package calculator provides functions for calculating toll fees based on vehicle type, dates, and tax rules.
package calculator

import (
	taxrules "congestion-calculator-manager/app/tax_rules"
	"congestion-calculator-manager/app/vehicles"
	"sort"
	"time"
)

// GetTax calculates the total toll fee for a vehicle based on given dates and tax rules.
// It handles time intervals and computes fees accordingly.
//
// Parameters:
//   - vehicle: The vehicle for which to calculate the toll fee.
//   - dates: Slice of time.Time representing the dates for toll calculation.
//   - taxRule: Pointer to a vehicles.TaxRule containing custom tax rules (can be nil for default rules).
//
// Returns:
//   - int: Total toll fee for the provided vehicle and dates.
func GetTax(vehicle vehicles.Vehicle, dates []time.Time, isCustom bool, taxRule taxrules.TaxRule) int {
	if len(dates) == 0 {
		return 0
	}

	// first we must sort slice
	sort.Slice(dates, func(i, j int) bool {
		return dates[i].Before(dates[j])
	})

	intervalStart := dates[0]
	totalFee := 0
	highestFee := 0

	for _, date := range dates {
		var nextFee int
		if isCustom {
			nextFee = getTollCustomFee(date, vehicle, taxRule)
		} else {
			nextFee = getTollFee(date, vehicle)
		}

		diffInMinutes := date.Sub(intervalStart).Minutes()

		if diffInMinutes <= 60 {
			if nextFee > highestFee {
				highestFee = nextFee
			}
		} else {
			totalFee += highestFee
			intervalStart = date
			highestFee = nextFee
		}
	}

	// Add the last highest fee for the 60-minute window
	// that should be covering both cases if last one is within or more then 60min
	totalFee += highestFee

	// Ensure totalFee does not exceed 60
	if totalFee > 60 {
		totalFee = 60
	}

	return totalFee
}

// isTollFreeVehicle checks if a vehicle is toll-free based on its tax exclusion status.
//
// Parameters:
//   - v: The vehicle to check.
//
// Returns:
//   - bool: True if the vehicle is toll-free, false otherwise.
func IsTollFreeVehicle(v vehicles.Vehicle) bool {
	if v == nil {
		return false
	}
	return v.IsTaxExcluded()
}

// getTollFee computes the toll fee based on specific time intervals for non-custom tax rules.
//
// Parameters:
//   - t: The time for which to calculate the toll fee.
//   - v: The vehicle for which to calculate the toll fee.
//
// Returns:
//   - int: The toll fee for the provided time and vehicle.
func getTollFee(t time.Time, v vehicles.Vehicle) int {
	if IsTollFreeDate(t) || IsTollFreeVehicle(v) {
		return 0
	}

	hour, minute := t.Hour(), t.Minute()

	if hour == 6 && minute >= 0 && minute <= 29 {
		return 8
	}
	if hour == 6 && minute >= 30 && minute <= 59 {
		return 13
	}
	if hour == 7 && minute >= 0 && minute <= 59 {
		return 18
	}
	if hour == 8 && minute >= 0 && minute <= 29 {
		return 13
	}
	if hour >= 8 && hour <= 14 && minute >= 30 && minute <= 59 {
		return 8
	}
	if hour == 15 && minute >= 0 && minute <= 29 {
		return 13
	}
	if hour == 15 && minute >= 0 || hour == 16 && minute <= 59 {
		return 18
	}
	if hour == 17 && minute >= 0 && minute <= 59 {
		return 13
	}
	if hour == 18 && minute >= 0 && minute <= 29 {
		return 8
	}

	return 0
}

// getTollCustomFee computes the toll fee based on custom tax rules.
//
// Parameters:
//   - t: The time for which to calculate the toll fee.
//   - v: The vehicle for which to calculate the toll fee.
//   - taxRules: The custom tax rules to apply.
//
// Returns:
//   - int: The toll fee for the provided time, vehicle, and custom tax rules.
func getTollCustomFee(t time.Time, v vehicles.Vehicle, taxRules taxrules.TaxRule) int {
	if IsTollFreeDate(t) || IsToolFreeDateWithCustomRules(t, taxRules) {
		return 0
	}
	hour := t.Hour()

	//Here we could add minutes,seconds it is same logic
	//this is just for demonstration
	for _, taxFeeRule := range taxRules.HourlyPrices {
		if hour >= taxFeeRule.StartHour && hour <= taxFeeRule.EndHour {
			return taxFeeRule.Rate
		}
	}

	return 0
}

// isTollFreeDate checks if a given date is toll-free based on predefined conditions.
//
// Parameters:
//   - date: The date to check.
//
// Returns:
//   - bool: True if the date is toll-free, false otherwise.
func IsTollFreeDate(date time.Time) bool {
	year := date.Year()
	month := date.Month()
	day := date.Day()

	if date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
		return true
	}

	if year == 2013 {
		if month == 1 && day == 1 ||
			month == 3 && (day == 28 || day == 29) ||
			month == 4 && (day == 1 || day == 30) ||
			month == 5 && (day == 1 || day == 8 || day == 9) ||
			month == 6 && (day == 5 || day == 6 || day == 21) ||
			month == 7 ||
			month == 11 && day == 1 ||
			month == 12 && (day == 24 || day == 25 || day == 26 || day == 31) {
			return true
		}
	}
	return false
}

// isToolFreeDateWithCustomRules checks if a given date is toll-free based on custom tax rules.
//
// Parameters:
//   - date: The date to check.
//   - taxRules: The custom tax rules to apply.
//
// Returns:
//   - bool: True if the date is toll-free based on custom rules, false otherwise.
func IsToolFreeDateWithCustomRules(date time.Time, taxRules taxrules.TaxRule) bool {
	year := date.Year()
	month := date.Month()
	day := date.Day()

	if !taxRules.TaxOnWeekend && date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
		return true
	}

	for _, exMonth := range taxRules.ExcludedMonths {
		if exMonth == int(month) {
			return true
		}
	}
	for _, exDay := range taxRules.ExcludedDays {
		if exDay == day {
			return true
		}
	}

	for _, exDates := range taxRules.ExcludedDates {
		if exDates.Year() == year && exDates.Month() == month && exDates.Day() == day {
			return true
		}
	}
	return false
}
