package main

import "fmt"

func main() {

	number1 := 12
	number2 := 20

	// outer if statement
	if number1 >= number2 {

		// inner if statement
		if number1 == number2 {
			fmt.Printf("Result: %d == %d", number1, number2)
			// inner else statement
		} else {
			fmt.Printf("Result: %d > %d", number1, number2)
		}

		// outer else statement
	} else {
		fmt.Printf("Result: %d < %d", number1, number2)
	}
}
