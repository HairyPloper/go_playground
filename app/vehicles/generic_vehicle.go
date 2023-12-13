// Package vehicles provides a generic representation of vehicles for congestion tax calculation.
package vehicles

import (
	"time"
)

// GenericVehicle represents the structure for a generic vehicle, including license plate, type, tax exclusion status, and times.
type GenericVehicle struct {
	LicensePlate string      `json:"license_plate"`
	Type         string      `json:"type"`
	TaxExcluded  bool        `json:"tax_excluded"`
	Times        []time.Time `json:"times"`
}

// GetVehicleType returns the type of the generic vehicle.
func (gv GenericVehicle) GetVehicleType() string {
	return gv.Type
}

// IsValidLicensePlate checks if the license plate of the generic vehicle is valid.
func (gv GenericVehicle) IsValidLicensePlate() bool {
	return gv.LicensePlate != ""
}

// IsTaxExcluded checks if the generic vehicle is tax excluded.
func (gv GenericVehicle) IsTaxExcluded() bool {
	return gv.TaxExcluded
}
