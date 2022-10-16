// Program to create a slice from an array

package main

import "fmt"

func main() {

	// an integer array
	numbers := [8]int{10, 20, 30, 40, 50, 60, 70, 80}

	// create slice from an array
	sliceNumbers := numbers[4:7]

	fmt.Println(sliceNumbers)

}
