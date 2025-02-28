package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func (c *Circle) setRadius(newRadius float64) {
	c.radius = newRadius
	// fmt.Println("mudando o valor do raio", c.radius, c.area())
}

func main() {
	c1 := Circle{3}

	fmt.Println(c1.area())
	fmt.Println(c1.perimeter())

	fmt.Println("setando o raio")
	c1.setRadius(4)

	fmt.Println(c1.area())
	fmt.Println(c1.perimeter())
}
