// Program to convert a string to uppercaseand lowercase

package main

import (
	"fmt"
	"strings"
)

func main() {

	text1 := "go is fun to learn"

	// convert to uppercase
	text1 = strings.ToUpper(text1)

	fmt.Println(text1)

	text2 := "I LOVE GOLANG"

	// convert to lowercase
	text2 = strings.ToLower(text2)
	fmt.Println(text2)
}
