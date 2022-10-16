// Program to pass pointer as a function argument

package main

import "fmt"

// function definition with a pointer argument
func update(num *int) {

	// dereference the pointer
	*num = 30

}

func main() {

	var number = 55

	// function call
	update(&number)

	fmt.Println("The number is", number)

}
