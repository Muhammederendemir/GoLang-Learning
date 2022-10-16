// Program to add elements to a slice

package main

import "fmt"

func main() {
	primeNumbers := []int{2, 3}

	// add elements 5, 7 to the slice
	primeNumbers = append(primeNumbers, 5, 7)
	fmt.Println("Prime Numbers:", primeNumbers)
}
