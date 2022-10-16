// Program to return the area of a rectangle

package main

import "fmt"

func main() {

	// anonymous function
	area := func(length, breadth int) int {
		return length * breadth
	}

	// function call using variable name
	fmt.Println("The area of rectangle is", area(3, 4))

}
