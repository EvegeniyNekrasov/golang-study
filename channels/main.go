package main

import (
	"fmt"
	"time"
)


func printMessage(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println(text)
		time.Sleep(800 * time.Millisecond)
	}
	channel <- "Done!"
}


func main() {
	channel := make(chan string)
	go printMessage("Go is great", channel)
	go printMessage("Other message", channel)
	response := <- channel
	fmt.Println(response)
	
	buffer := make(chan int, 2)

	buffer <- 1
	buffer <- 2

	fmt.Println(<- buffer)
	fmt.Println(<- buffer)
}