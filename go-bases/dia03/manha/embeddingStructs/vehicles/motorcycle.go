package vehicles

import "fmt"

type Motorcycle struct {
	Engine
	Chassis
	Bodywork
}

func (m Motorcycle) Start() {
	fmt.Println("The motorcycle is starting with a vroom!")
}
func (m Motorcycle) DisplayInfo() {
	fmt.Printf("This motorcycle has a %s body, a %s chassis, and a %d horsepower %s engine.\n",
		m.Bodywork.Color, m.Chassis.Material, m.Engine.Horsepower, m.Engine.Type)
}
