package main

import (
	"fmt"
	"time"
)

func main() {
	requests := make(chan int)

	for i := range 1_000 {
		go processRequest(i, &requests)
	}

	for i := range 100_000 {
		requests <- i
	}
}

func processRequest(id int, channel *chan int) {
	for message := range *channel {
		fmt.Printf("processing message from worker %d, message %d\n", id, message)
		time.Sleep(time.Millisecond * 200)
	}
}
