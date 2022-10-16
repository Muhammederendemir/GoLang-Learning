package main

import "fmt"

func main() {

	// create a map
	personAge := map[string]int{"Hermione": 21, "Harry": 20, "John": 25}
	fmt.Println("Initial Map: ", personAge)

	// remove element of map with key John
	delete(personAge, "John")

	fmt.Println("Updated Map: ", personAge)
}
