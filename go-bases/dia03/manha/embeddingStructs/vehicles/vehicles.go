package vehicles

type Engine struct {
	Horsepower int
	Type       string
}

type Chassis struct {
	Material string
}

type Bodywork struct {
	Color string
}

type Vehicle interface {
	Start()
	DisplayInfo()
}
