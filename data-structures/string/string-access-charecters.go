// Program to access individual character of string

package main

import "fmt"

func main() {

	// create and initialize a string
	name := "Programiz"

	// access first character
	fmt.Printf("%c\n", name[0]) // P

	// access fourth character
	fmt.Printf("%c\n", name[3]) // g

	// access last character
	fmt.Printf("%c", name[8]) // z
}
