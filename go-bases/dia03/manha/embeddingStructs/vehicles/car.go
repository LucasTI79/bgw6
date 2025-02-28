package vehicles

import "fmt"

type Car struct {
	Engine
	Chassis
	Bodywork
	NumberOfDoors int
}

func (c Car) Start() {
	fmt.Println("The car is starting with a roar!")
}
func (c Car) DisplayInfo() {
	fmt.Printf("This car has %d doors, a %s body, a %s chassis, and a %d horsepower %s engine.\n",
		c.NumberOfDoors, c.Bodywork.Color, c.Chassis.Material, c.Engine.Horsepower, c.Engine.Type)
}
