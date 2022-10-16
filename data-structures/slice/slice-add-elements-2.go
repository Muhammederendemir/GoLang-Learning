// Program to add elements of one slice to another

package main

import "fmt"

func main() {

	// create two slices
	evenNumbers := []int{2, 4}
	oddNumbers := []int{1, 3}

	// add elements of oddNumbers to evenNumbers
	evenNumbers = append(evenNumbers, oddNumbers...)
	fmt.Println("Numbers:", evenNumbers)
}
