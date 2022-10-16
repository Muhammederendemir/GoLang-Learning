package main

import "fmt"

// empty interface as function parameter
func displayValue(i interface{}) {
	fmt.Println(i)
}

func main() {

	a := "Muhammed Eren Demir"
	b := 20
	c := false

	// pass string to the function
	displayValue(a)

	// pass integer number to the function
	displayValue(b)

	// pass boolean value to the function
	displayValue(c)

}
