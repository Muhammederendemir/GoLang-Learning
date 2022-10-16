// Program to get the value pointed by a pointer

package main

import "fmt"

func main() {

	var name = "Muhammed Eren"
	var ptr *string

	ptr = &name

	// * to get the value pointed by ptr
	fmt.Println(*ptr) // Muhammed Eren

}
