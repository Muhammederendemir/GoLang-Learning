// Program to return a pointer from a function

package main

import "fmt"

func main() {

	// function call
	result := display()
	fmt.Println("I am ", *result)

}

func display() *string {

	message := "Muhammed Eren"

	// returns the address of message
	return &message

}
