package main

import (
	"fmt"
	"time"
)

func printMessage(text string, channel chan string) {
	for i := 0; i < 10; i++ {
		fmt.Println(text)
		time.Sleep(time.Second)
	}
	channel <- "Done!"
}

func main() {
	channel := make(chan string)

	go printMessage("Go Routines", channel)
	//	printMessage("Heyyyyy")
	println(<-channel)
}
