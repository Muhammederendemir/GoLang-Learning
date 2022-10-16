// Program to display a number if it is positive

package main

import "fmt"

func main() {
	number := 15

	// true if number is less than 0
	if number > 0 {
		fmt.Printf("%d is a positive number\n", number)
	}

	fmt.Println("Out of the loop")
}
