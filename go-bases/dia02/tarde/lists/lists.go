package lists

import "fmt"

func ManipularArray() {
	//0 , 1 , 2 , 3 , 4
	array := [5]int{1, 2, 3, 4, 5}
	array2 := [10]*int{}

	// fmt.Println("Array inicial:", array)
	fmt.Println("Array inicial:", array2)

	array[0] = 2
	// fmt.Println("Array alterado na posilção 0:", array)
	fmt.Println(&array)
}
