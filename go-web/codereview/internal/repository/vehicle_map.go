package repository

import (
	"app/internal"
	"strings"
)

// NewVehicleMap is a function that returns a new instance of VehicleMap
func NewVehicleMap(db map[int]internal.Vehicle) *VehicleMap {
	// default db
	defaultDb := make(map[int]internal.Vehicle)
	if db != nil {
		defaultDb = db
	}
	return &VehicleMap{db: defaultDb}
}

// VehicleMap is a struct that represents a vehicle repository
type VehicleMap struct {
	// db is a map of vehicles
	db map[int]internal.Vehicle
}

// FindAll is a method that returns a map of all vehicles
func (repository *VehicleMap) FindAll() (vehicles map[int]internal.Vehicle, err error) {
	vehicles = make(map[int]internal.Vehicle)

	// copy db
	for id, vehicle := range repository.db {
		vehicles[id] = vehicle
	}

	return
}

func (
	repository *VehicleMap,
) FindByColorAndYear(
	color string,
	year int,
) (
	vehicles map[int]internal.Vehicle,
	err error,
) {
	vehicles = make(map[int]internal.Vehicle)

	for id, vehicle := range repository.db {
		if strings.EqualFold(vehicle.Color, color) && vehicle.FabricationYear == year {
			vehicles[id] = vehicle
		}
	}

	return
}
