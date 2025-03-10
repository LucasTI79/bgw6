package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// PascalCase => structs go
// camelCase  => json
// snake_case => bd

type product struct {
	Name      string `json:"name"`
	Price     int
	Published bool
}

func main() {

	// Exemplo Marshal (processo de encode)
	p1 := product{
		Name:      "MacBook Pro",
		Price:     1500,
		Published: true,
	}

	jsonData, err := json.Marshal(p1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))

	// Exemplo Unmarshal (processo de decode)
	data := `{"Name": "MacBook Air", "Price": 900, "Published": true}`

	var p product

	if err := json.Unmarshal([]byte(data), &p); err != nil {
		log.Fatal(err)
	}

	fmt.Println(p)
}
