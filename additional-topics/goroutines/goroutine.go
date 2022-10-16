package main

import (
	"fmt"
)

// create a function
func display(message string) {

	fmt.Println(message)
}

func main() {

	// call goroutine
	go display("Process 1")

	display("Process 2")

}
