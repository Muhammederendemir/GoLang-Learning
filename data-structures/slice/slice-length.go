// Program to find the length of a slice

package main

import "fmt"

func main() {

	// create a slice of numbers
	numbers := []int{1, 5, 8, 0, 3}

	// find the length of the slice
	length := len(numbers)

	fmt.Println("Length:", length)

}
