package main

import "fmt"

func divide(num1, num2 int) error {

	// returns error
	if num2 == 0 {
		return fmt.Errorf("%d / %d\nCannot Divide a Number by zero", num1, num2)
	}

	// returns the result of division
	return nil
}

func main() {

	err := divide(4, 0)

	// error found
	if err != nil {
		fmt.Printf("error: %s", err)

		// error not found
	} else {
		fmt.Println("Valid Division")
	}
}
