// Program to return an anonymous function

package main

import "fmt"

// function that returns an anonymous function
func displayNumber() func() int {

	number := 10
	return func() int {
		number++
		return number
	}
}

func main() {

	a := displayNumber()

	fmt.Println(a())

}
