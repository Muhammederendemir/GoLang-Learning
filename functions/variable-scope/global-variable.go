// Program to illustrate global variable

package main

import "fmt"

// declare global variable before main function
var sum int

func addNumbers() {

	// local variable
	sum = 9 + 5
}

func main() {

	addNumbers()

	// can access sum
	fmt.Println("Sum is", sum)

}
