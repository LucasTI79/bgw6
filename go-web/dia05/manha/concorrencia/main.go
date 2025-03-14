package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int)

	for i := 0; i < 1_000; i++ {
		go processRequest(i, &requests)
	}

	for i := 0; i < 100_000; i++ {
		requests <- i
	}
}

func processRequest(id int, channel *chan int) {
	for message := range *channel {
		fmt.Printf("processing message from worker %d, message %d\n", id, message)
		time.Sleep(time.Millisecond * 200)
	}
}
