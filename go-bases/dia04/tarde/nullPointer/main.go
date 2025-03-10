package main

import (
	"fmt"
)

/*
	Equivalente ao:

	null pointer exception
	ou
	cannot read properties of undefined
*/

type Dog struct {
	Name string
}

func (s *Dog) WoofWoof() {
	fmt.Println(s.Name, "Goes woof woof")
}
func main() {
	s := &Dog{"Sammy"}
	s = nil

	s.WoofWoof()
}

// 0 -> sucesso
// 1 -> 125 (erro)
// exit code
