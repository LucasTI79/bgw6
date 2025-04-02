package service

import "app/internal"

// NewVehicleDefault is a function that returns a new instance of VehicleDefault
func NewVehicleDefault(rp internal.VehicleRepository) *VehicleDefault {
	return &VehicleDefault{repository: rp}
}

// VehicleDefault is a struct that represents the default service for vehicles
type VehicleDefault struct {
	// rp is the repository that will be used by the service
	repository internal.VehicleRepository
}

// FindAll is a method that returns a map of all vehicles
func (service *VehicleDefault) FindAll() (v map[int]internal.Vehicle, err error) {
	v, err = service.repository.FindAll()
	return
}

func (service *VehicleDefault) FindByColorAndYear(color string, year int) (vehicles map[int]internal.Vehicle, err error) {
	vehicles, err = service.repository.FindByColorAndYear(color, year)
	return
}
