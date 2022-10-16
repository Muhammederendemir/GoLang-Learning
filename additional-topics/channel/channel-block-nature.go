package main

import "fmt"

func main() {

	// create channel
	ch := make(chan string)

	// function call with goroutine
	go sendData(ch)

	// receive channel data
	fmt.Println(<-ch)

}

func sendData(ch chan string) {

	// data sent to the channel
	ch <- "Received. Send Operation Successful"
	fmt.Println("No receiver! Send Operation Blocked")

}
