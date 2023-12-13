// Package helpers contains all common functions used in other packages
package helpers

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// ReadCityDataFromFile reads city tax data from a JSON file located in the "cities" directory.
//
// Parameters:
//   - cityName: The name of the city to read data for.
//
// Returns:
//   - string: content of json file read from this
//   - err: An error if any occurs during the file reading process.
func ReadContentFromJsonFile(cityName string) (string, error) {

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return "", err
	}

	// Construct the file path for the specified city
	filePath := filepath.Join(wd, "cities", fmt.Sprintf("%s.json", strings.ToLower(cityName)))

	// Read the JSON file
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("data for specified city does not exist on disk", err)
		return "", err
	}

	return string(fileContent), nil
}

// generateRandomDate generates a random time within the specified year.
// It takes the year as an input parameter and returns a random time.Time within that year.
//
// Parameters:
//   - year: The year for which to generate a random date.
//
// Returns:
//   - time.Time: A randomly generated time within the specified year.
func GenerateRandomDate(year int) time.Time {
	start := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(year, time.December, 31, 0, 0, 0, 0, time.UTC)
	delta := end.Sub(start)
	randomSeconds := rand.Int63n(int64(delta.Seconds()))
	randomDuration := time.Duration(randomSeconds) * time.Second

	return start.Add(randomDuration)
}

// generateNumberOfDates generates a specified number of random dates within the given year.
// It takes the number of dates to generate (`numDates`) and the year as input parameters.
// It returns a slice of time.Time containing the randomly generated dates.
//
// Parameters:
//   - numDates: The number of random dates to generate.
//   - year: The year for which to generate random dates.
//
// Returns:
//   - []time.Time: A slice containing the randomly generated dates within the specified year.

func GenerateNumberOfDates(numDates int, year int) (randomDates []time.Time) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numDates; i++ {
		randomDate := GenerateRandomDate(year)
		//fmt.Println(randomDate)
		randomDates = append(randomDates, randomDate)
	}
	return randomDates
}
