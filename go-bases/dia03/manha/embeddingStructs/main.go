package main

import "github.com/bgw6/embedding/vehicles"

func main() {
	var vehicle vehicles.Vehicle
	vehicle = vehicles.Motorcycle{}
	vehicle.Start()
	vehicle.DisplayInfo()
	vehicle = vehicles.Car{}
	vehicle.Start()
	vehicle.DisplayInfo()
}
