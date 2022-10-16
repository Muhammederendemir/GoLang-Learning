package main

import "fmt"

func displayNumbers() func() int {
	number := 0

	// inner function
	return func() int {
		number++
		return number
	}

}

func main() {

	// returns a closure
	num1 := displayNumbers()

	fmt.Println(num1())
	fmt.Println(num1())
	fmt.Println(num1())

	// returns a new closure
	num2 := displayNumbers()
	fmt.Println(num2())
	fmt.Println(num2())

}
