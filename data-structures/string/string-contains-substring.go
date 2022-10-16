// Program to illustrate the use of Contains()

package main

import (
	"fmt"
	"strings"
)

func main() {

	text := "Go Programming"
	substring1 := "Go"
	substring2 := "Golang"

	// check if Go is present in Go Programming
	result := strings.Contains(text, substring1)
	fmt.Println(result)

	// check if Golang is present in Go Programming
	result = strings.Contains(text, substring2)
	fmt.Println(result)
}
