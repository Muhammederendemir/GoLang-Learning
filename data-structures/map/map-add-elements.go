package main

import "fmt"

func main() {

	// create a map
	students := map[int]string{1: "John", 2: "Lily"}
	fmt.Println("Initial Map: ", students)

	// add element with key 3
	students[3] = "Robin"

	// add element with key 5
	students[5] = "Julie"

	fmt.Println("Updated Map: ", students)
}
