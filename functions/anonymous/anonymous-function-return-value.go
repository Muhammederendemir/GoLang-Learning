// Program to return value from an anonymous function

package main

import "fmt"

func main() {

	// anonymous function
	var sum = func(n1, n2 int) int {
		sum := n1 + n2

		return sum
	}

	// function call
	result := sum(5, 3)

	fmt.Println("Sum is:", result)

}
