package main

import "fmt"

func main() {

	// anonymous function
	var greet = func() {
		fmt.Println("Hello, how are you")
	}

	// function call
	greet()

}
