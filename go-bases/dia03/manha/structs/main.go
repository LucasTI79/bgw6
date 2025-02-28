package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name       string  `json:"name"`
	Gender     string  `json:"gender"`
	Age        int     `json:"age"`
	Profession string  `json:"profession"`
	Weight     float64 `json:"weight"`
	// LastName   string  `json:"lastName",bd:"last_name"`
	Likes Preferences
}

type Preferences struct {
	Foods  string
	Movies string
	Series string
	Animes string
	Sports string
}

func main() {
	p1 := Person{"Celeste", "Woman", 34, "Engineer", 65.5, Preferences{"chicken", "titanic", "", "", ""}}

	p1.Likes.Animes = "spy family"

	fmt.Printf("%+v\n", p1)

	json, _ := json.Marshal(p1)
	fmt.Println(string(json))
}
