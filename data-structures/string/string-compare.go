// Program to compare string using Compare()

package main

import (
	"fmt"
	"strings"
)

func main() {

	// create three strings
	string1 := "Programiz"
	string2 := "Programiz Pro"
	string3 := "Programiz"

	// compare strings
	fmt.Println(strings.Compare(string1, string2)) // -1
	fmt.Println(strings.Compare(string2, string3)) // 1
	fmt.Println(strings.Compare(string1, string3)) // 0

}
