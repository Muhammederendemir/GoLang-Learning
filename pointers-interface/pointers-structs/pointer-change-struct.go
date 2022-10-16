// Program to change the struct member using pointer

package main

import "fmt"

// create struct
type Weather struct {
	city        string
	temperature int
}

func main() {

	// create struct variable
	weather := Weather{"California", 20}
	fmt.Println("Initial Weather:", weather)

	// create struct type pointer
	ptr := &weather

	// change value of temperature to 25
	ptr.temperature = 25

	fmt.Println("Updated Weather:", weather)

}
