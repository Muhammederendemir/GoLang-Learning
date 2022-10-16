package main

import "fmt"

func main() {

	// create an array
	var arrayOfIntegers = [...]int{1, 5, 8, 0, 3, 10}

	// find the length of array using len()
	length := len(arrayOfIntegers)

	fmt.Println("The length of array is", length)
}
