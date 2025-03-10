package main

import "fmt"

func funcaoQualquer() {
	fmt.Println("This function is executed despite a panic occurring")
}

func main() {
	//apply “defer” to the invocation of an anonymous function
	defer funcaoQualquer()
	//create a panic with a message that it occurred
	panic("panic occurred!!!")
}
