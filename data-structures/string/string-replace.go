// Program using Replace() to replace strings

package main

import (
	"fmt"
	"strings"
)

func main() {

	text := "car"
	fmt.Println("Old String:", text)

	// replace r with t
	replacedText := strings.Replace(text, "r", "t", 1)

	fmt.Println("New String:", replacedText)
}
