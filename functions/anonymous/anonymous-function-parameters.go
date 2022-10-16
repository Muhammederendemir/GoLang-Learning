// Program to pass arguments in an anonymous function

package main

import "fmt"

func main() {

	// anonymous function with arguments
	var sum = func(n1, n2 int) {
		sum := n1 + n2
		fmt.Println("Sum is:", sum)
	}

	// function call
	sum(5, 3)

}
