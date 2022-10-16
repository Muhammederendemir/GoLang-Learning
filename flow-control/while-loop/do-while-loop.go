// Program to print number from 1 to 5

package main

import "fmt"

func main() {
	number := 1

	// loop that runs infinitely
	for {

		// condition to terminate the loop
		if number > 5 {
			break
		}

		fmt.Printf("%d\n", number)
		number++

	}
}
