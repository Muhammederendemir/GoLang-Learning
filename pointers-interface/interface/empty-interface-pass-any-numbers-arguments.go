package main

import "fmt"

// function with an empty interface as its parameter
func displayValue(i ...interface{}) {
	fmt.Println(i)
}

func main() {

	a := "I am Muhammed Eren Demir"
	b := 20
	c := false

	// function call with a single parameter
	displayValue(a)

	// function call with 2 parameters
	displayValue(a, b)

	// function call with 3 parameters
	displayValue(a, b, c)

}
