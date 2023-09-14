package service

import (
	"app/internal/domain"
	"errors"
)

// ServiceVehicle is the interface that wraps the basic methods for a vehicle service.
// - conections with external apis
// - business logic
type ServiceVehicle interface {
	// GetAll returns all vehicles
	GetAll() (v []*domain.Vehicle, err error)
	SaveVehicles(vehiclesList []domain.Vehicle) (err error)
	GetByColorAndYear(color string, year int) (vehiclesList []*domain.Vehicle, err error)
}

var (
	// ErrServiceVehicleInternal is returned when an internal error occurs.
	ErrServiceVehicleInternal = errors.New("service: internal error")

	// ErrServiceVehicleNotFound is returned when no vehicle is found.
	ErrServiceVehicleNotFound = errors.New("service: vehicle not found")

	// ErrServiceVehiclesNotFound is returned when no vehicles is found.
	ErrServiceVehiclesNotFound = errors.New("service: vehicles not found")

	// ErrServiceVehicleAlreadyExist is returned when a vehicle already exists in database
	ErrServiceVehicleAlreadyExist = errors.New("service: vehicle already exists")
)
