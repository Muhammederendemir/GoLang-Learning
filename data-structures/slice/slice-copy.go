// Program to copy one slice to another

package main

import "fmt"

func main() {

	// create two slices
	primeNumbers := []int{2, 3, 5, 7}
	numbers := []int{1, 2, 3}

	// copy elements of primeNumbers to numbers
	copy(numbers, primeNumbers)

	// print numbers
	fmt.Println("Numbers:", numbers)
}
