package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Starting... ")
	_, err := os.Open("no-file.txt")
	// defer file.Close()
	if err != nil {
		// log.Fatalf("ocorreu um erro doidao: %s", err)
		panic(err)
	}
	fmt.Println("End")
}
