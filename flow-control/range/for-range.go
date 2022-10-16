// Program using range with array

package main

import "fmt"

func main() {

	// array of numbers
	numbers := [5]int{21, 24, 27, 30, 33}

	// use range to iterate over the elements of array
	for index, item := range numbers {
		fmt.Printf("numbers[%d] = %d \n", index, item)
	}

}
