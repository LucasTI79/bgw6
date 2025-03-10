package loops

import "fmt"

func LoopInfinito() {
	sum := 0
	for {
		fmt.Println(sum)
		sum++ // sum + 1
	}
}

func LoopWhile() {
	sum := 1
	for sum < 10 {
		sum += sum
	}

	fmt.Println(sum)
}

func LoopComContinue() {
	for i := 0; i <= 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}

func LoopComBreak() {
	for i := 0; i <= 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}
}
