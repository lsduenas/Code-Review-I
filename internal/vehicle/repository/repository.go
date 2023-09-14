package repository

import (
	"app/internal/domain"
	"errors"
)

// RepositoryVehicle is the interface that wraps the basic methods for a vehicle repository.
type RepositoryVehicle interface {
	// GetAll returns all vehicles
	GetAll() (v []*domain.Vehicle, err error)
	WriteJSONFile(vehiclesList []* VehicleAttributesJSON) (err error)

	// Save() (err error)
	// GetByColorAndYear(color string, year string) (v []*domain.Vehicle, err error)
	// GetByBrandAndYearsRange(brand string, start_year int, end_year int) (v []*domain.Vehicle, err error)
	// GetAverageSpeedByBrand(brand string) (average float64)
	SaveVehicles(vehicleList []domain.Vehicle) (err error)
	// UpdateMaxSpeed(id int, speed int) (err error)
	// GetByFuelType(fuel_type string) (v []*domain.Vehicle, err error)
	// DeleteById(id int) (err error)
	// GetByTransmissionType(transmission_type string) (v []*domain.Vehicle, err error)
	// UpdateFuelType(id int, fuel_type string) (err error)
	// GetAverageCapacityByBranch(branch string) (average float64) // capacity -> passangers
	// GetByDimensions(min_height float64, max_heigth float64, min_width float64, max_width float64) (v []*domain.Vehicle, err error)
	// GetByWeight(min_weight float64, max_weight float64) (v []*domain.Vehicle, err error) // query params
}

var (
	// ErrRepositoryVehicleInternal is returned when an internal error occurs.
	ErrRepositoryVehicleInternal = errors.New("repository: internal error")

	// ErrRepositoryVehicleNotFound is returned when a vehicle is not found.
	ErrRepositoryVehicleNotFound = errors.New("repository: vehicle not found")

	// ErrRepositoryCantUpdateJSONFile is returned when it's not posible update the JSON file 
	ErrRepositoryVehicleCantUpdateJSONFile = errors.New("repository: can't update JSON file")


)
