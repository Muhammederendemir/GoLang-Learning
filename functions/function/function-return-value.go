package main

import "fmt"

// function definition
func addNumbers(n1 int, n2 int) int {
	sum := n1 + n2
	return sum
}

func main() {

	// function call
	result := addNumbers(21, 13)

	fmt.Println("Sum:", result)
}
