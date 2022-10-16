package main

import "fmt"

var sum = 0

// regular function to calculate square of numbers
func findSquare(num int) int {
	square := num * num
	return square
}

func main() {

	// anonymous function that returns sum of numbers
	sum := func(number1 int, number2 int) int {
		return number1 + number2
	}

	// function call
	result := findSquare(sum(6, 9))
	fmt.Println("Result is:", result)

}
