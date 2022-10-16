package main

import "fmt"

// function to add two numbers
func addNumbers() {
	n1 := 12
	n2 := 8

	sum := n1 + n2
	fmt.Println("Sum:", sum)
}

func main() {
	// function call
	addNumbers()
}
