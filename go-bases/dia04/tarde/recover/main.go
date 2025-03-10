package main

import (
	"fmt"
)

func isPair(num int) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	if (num % 2) != 0 {
		panic("Not an even number")
	}
	fmt.Println(num, "is an even number!")
}

func main() {
	num := 3
	isPair(num)
	fmt.Println("Execution completed!")

	// http.ListenAndServe(":8080", nil)
}

/*
    Print
	Println
	Printf
	Sprintf
	Errorf
*/
