package repository

import (
	"app/internal/domain"
	"encoding/json"
	"errors"
	"os"
)

func NewRepositoryVehicleInMemory(db map[int]*domain.VehicleAttributes, path string) *RepositoryVehicleInMemory {
	return &RepositoryVehicleInMemory{db: db, path: path}
}

// RepositoryVehicleInMemory is an struct that represents a vehicle storage in memory.
type RepositoryVehicleInMemory struct {
	// db is the database of vehicles.
	db   map[int]*domain.VehicleAttributes
	path string
}

type VehicleAttributesJSON struct {
	Id           int
	Brand        string
	Model        string
	Registration string
	Year         int
	Color        string
	MaxSpeed     int
	FuelType     string
	Transmission string
	Passengers   int
	Height       float64
	Width        float64
	Weight       float64
}

// GetAll returns all vehicles
func (s *RepositoryVehicleInMemory) GetAll() (v []*domain.Vehicle, err error) {
	// check if the database is empty
	if len(s.db) == 0 {
		err = ErrRepositoryVehicleNotFound
		return
	}

	// get all vehicles from the database
	v = make([]*domain.Vehicle, 0, len(s.db))
	for key, value := range s.db {
		v = append(v, &domain.Vehicle{
			Id:         key,
			Attributes: *value,
		})
	}

	return
}

// Save vehicles list
func (s *RepositoryVehicleInMemory) SaveVehicles(vehicleList []domain.Vehicle) (err error) {
	// check if the database is empty
	if len(s.db) == 0 {
		err = ErrRepositoryVehicleNotFound
		return
	}

	// get all vehicles from the database
	v := make([]*VehicleAttributesJSON, 0, len(s.db))
	for key, value := range s.db {
		v = append(v, &VehicleAttributesJSON{
			Id:           key,
			Brand:        value.Brand,
			Model:        value.Model,
			Registration: value.Registration,
			Year:         value.Year,
			Color:        value.Color,
			MaxSpeed:     value.MaxSpeed,
			FuelType:     value.FuelType,
			Transmission: value.Transmission,
			Passengers:   value.Passengers,
			Height:       value.Height,
			Width:        value.Width,
			Weight:       value.Weight,
		})
	}

	// validar ID unico del vehicle
	for _, vehicle := range v {
		for _, newVehicle := range vehicleList {
			if vehicle.Id == newVehicle.Id {
				err = errors.New("Algún vehículo tiene un identificador ya existente.")
				return
			}
		}

	}

	// add new vehicles from list [] domain.Vehicle
	for _, value := range vehicleList {

		v = append(v, &VehicleAttributesJSON{
			Id:           value.Id,
			Brand:        value.Attributes.Brand,
			Model:        value.Attributes.Model,
			Registration: value.Attributes.Registration,
			Year:         value.Attributes.Year,
			Color:        value.Attributes.Color,
			MaxSpeed:     value.Attributes.MaxSpeed,
			FuelType:     value.Attributes.FuelType,
			Transmission: value.Attributes.Transmission,
			Passengers:   value.Attributes.Passengers,
			Height:       value.Attributes.Height,
			Width:        value.Attributes.Width,
			Weight:       value.Attributes.Weight,
		})
	}
	err = s.WriteJSONFile(v)
	if err != nil {
		err = errors.New("Internal server error")
		return
	}
	return
}

// update the json file of vehicles 
func (s *RepositoryVehicleInMemory) WriteJSONFile(vehiclesList []*VehicleAttributesJSON) (err error) {
	file, err := os.Create(s.path)
	if err != nil {
		err = errors.New("Can not create file")
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)                    // crea el encoder
	if err = encoder.Encode(vehiclesList); err != nil { // a través del encoder convierte la lista de structs en el json file
		err = ErrRepositoryVehicleCantUpdateJSONFile
		return
	}
	return nil
}
