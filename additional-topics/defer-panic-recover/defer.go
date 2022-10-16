package main

import "fmt"

func main() {

	// defer the execution of Println() function
	defer fmt.Println("Three")

	fmt.Println("One")
	fmt.Println("Two")

}
