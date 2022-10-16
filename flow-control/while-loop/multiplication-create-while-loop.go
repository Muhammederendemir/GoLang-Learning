// Program to create a multiplication table of 5 using while loop

package main

import (
	"fmt"
)

func main() {
	multiplier := 1

	// run while loop for 10 times
	for multiplier <= 10 {

		// find the product
		product := 5 * multiplier

		// print the multiplication table in format 5 * 1 = 5
		fmt.Printf("5 * %d = %d\n", multiplier, product)
		multiplier++
	}

}
