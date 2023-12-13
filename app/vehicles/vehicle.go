// Package vehicles provides a set of vehicle types and an interface for interacting with them in the context of congestion tax calculation.

package vehicles

import "errors"

// Vehicle is an interface that defines the methods expected from a vehicle type.
type Vehicle interface {
	GetVehicleType() string
	IsValidLicensePlate() bool
	IsTaxExcluded() bool
}

// GetVehicle creates and returns a Vehicle object based on the provided type and license information.
func GetVehicle(vehicleType string, licenseInfo string) (Vehicle, error) {
	var vehicle Vehicle

	switch vehicleType {
	case "Car":
		vehicle = Car{LicensePlate: licenseInfo}
	case "Bus":
		vehicle = Bus{LicensePlate: licenseInfo}
	case "Motorbike":
		vehicle = Motorbike{LicensePlate: licenseInfo}
	case "Military":
		vehicle = Military{LicensePlate: licenseInfo}
	case "Diplomat":
		vehicle = Diplomat{LicensePlate: licenseInfo}
	case "Emergency":
		vehicle = Emergency{LicensePlate: licenseInfo}
	case "Foreign":
		vehicle = Foreign{LicensePlate: licenseInfo}
	default:
		// Default to an unknown vehicle type
		return nil, errors.New("unknown or missing vehicle type")
	}
	if licenseInfo == "" {
		return nil, errors.New("invalid license for specified vehicle")
	}

	return vehicle, nil
}

// Bus represents a bus vehicle type.
type Bus struct {
	LicensePlate string
}

// GetVehicleType returns the type of the bus.
func (b Bus) GetVehicleType() string {
	return "Bus"
}

// IsValidLicensePlate checks if the license plate is valid for the bus.
func (b Bus) IsValidLicensePlate() bool {
	return b.LicensePlate != ""
}

// IsTaxExcluded checks if the bus is tax excluded.
func (b Bus) IsTaxExcluded() bool {
	return true
}

// Motorbike represents a motorbike vehicle type.
type Motorbike struct {
	LicensePlate string
	Type         int
}

// GetVehicleType returns the type of the motorbike.
func (m Motorbike) GetVehicleType() string {
	return "Motorbike"
}

// IsValidLicensePlate checks if the license plate is valid for the motorbike.
func (m Motorbike) IsValidLicensePlate() bool {
	return m.LicensePlate != ""
}

// IsTaxExcluded checks if the motorbike is tax excluded.
func (m Motorbike) IsTaxExcluded() bool {
	return true
}

// Car represents a car vehicle type.
type Car struct {
	LicensePlate string
}

// GetVehicleType returns the type of the car.
func (c Car) GetVehicleType() string {
	return "Car"
}

// IsValidLicensePlate checks if the license plate is valid for the car.
func (c Car) IsValidLicensePlate() bool {
	return c.LicensePlate != ""
}

// IsTaxExcluded checks if the car is tax excluded.
func (c Car) IsTaxExcluded() bool {
	return false
}

// Military represents a military vehicle type.
type Military struct {
	LicensePlate string
}

// GetVehicleType returns the type of the military vehicle.
func (m Military) GetVehicleType() string {
	return "Military"
}

// IsValidLicensePlate checks if the license plate is valid for the military vehicle.
func (m Military) IsValidLicensePlate() bool {
	return m.LicensePlate != ""
}

// IsTaxExcluded checks if the military vehicle is tax excluded.
func (m Military) IsTaxExcluded() bool {
	return true
}

// Diplomat represents a diplomat vehicle type.
type Diplomat struct {
	LicensePlate string
}

// GetVehicleType returns the type of the diplomat vehicle.
func (d Diplomat) GetVehicleType() string {
	return "Diplomat"
}

// IsValidLicensePlate checks if the license plate is valid for the diplomat vehicle.
func (d Diplomat) IsValidLicensePlate() bool {
	return d.LicensePlate != ""
}

// IsTaxExcluded checks if the diplomat vehicle is tax excluded.
func (d Diplomat) IsTaxExcluded() bool {
	return true
}

// Emergency represents an emergency vehicle type.
type Emergency struct {
	LicensePlate string
}

// GetVehicleType returns the type of the emergency vehicle.
func (e Emergency) GetVehicleType() string {
	return "Emergency"
}

// IsValidLicensePlate checks if the license plate is valid for the emergency vehicle.
func (e Emergency) IsValidLicensePlate() bool {
	return e.LicensePlate != ""
}

// IsTaxExcluded checks if the emergency vehicle is tax excluded.
func (e Emergency) IsTaxExcluded() bool {
	return true
}

// Foreign represents a foreign vehicle type.
type Foreign struct {
	LicensePlate string
}

// GetVehicleType returns the type of the foreign vehicle.
func (f Foreign) GetVehicleType() string {
	return "Foreign"
}

// IsValidLicensePlate checks if the license plate is valid for the foreign vehicle.
func (f Foreign) IsValidLicensePlate() bool {
	return f.LicensePlate != ""
}

// IsTaxExcluded checks if the foreign vehicle is tax excluded.
func (f Foreign) IsTaxExcluded() bool {
	return true
}
