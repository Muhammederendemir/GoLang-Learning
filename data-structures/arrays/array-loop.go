package main

import "fmt"

func main() {
	age := [...]int{12, 4, 5}

	// loop through the array
	for i := 0; i < len(age); i++ {
		fmt.Println(age[i])
	}

}
