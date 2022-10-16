// Program to illustrate code reusability in function

package main

import "fmt"

// function definition
func getSquare(num int) {
	square := num * num
	fmt.Printf("Square of %d is %d\n", num, square)
}

// main function
func main() {

	// call function 3 times
	getSquare(3)
	getSquare(5)
	getSquare(10)

}
