package main

import "fmt"

func main() {
	number := 10

	// checks if number is greater than 0
	if number > 0 {
		fmt.Printf("%d is a positive number\n", number)
	} else {
		fmt.Printf("%d is a negative number\n", number)
	}
}
