// Program using range with string

package main

import "fmt"

func main() {
	string := "Golang"
	fmt.Println("Index: Character")

	// i access index of each character
	// item access each character
	for i, item := range string {
		fmt.Printf("%d= %c \n", i, item)
	}

}
