package main

import (
	"fmt"
	"strings"
)

func main() {
	var message = "I Love Golang"

	// split string from space " "
	splittedString := strings.Split(message, " ")

	fmt.Println(splittedString)
}

// Output: [I Love Golang]
