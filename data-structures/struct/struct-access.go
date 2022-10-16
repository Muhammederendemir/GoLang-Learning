// Program to access the individual elements of struct

package main

import "fmt"

func main() {

	// declare a struct
	type Rectangle struct {
		length  int
		breadth int
	}

	// declare instance rect1 and defining the struct
	rect := Rectangle{22, 12}

	// access the length of the struct
	fmt.Println("Length:", rect.length)

	// access the breadth of the struct
	fmt.Println("Breadth:", rect.breadth)

	area := rect.length * rect.breadth
	fmt.Println("Area:", area)

}
