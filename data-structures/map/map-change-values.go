package main

import "fmt"

func main() {

	// create a map
	capital := map[string]string{"Nepal": "Kathmandu", "US": "New York"}
	fmt.Println("Initial Map: ", capital)

	// change value of US to Washington DC
	capital["US"] = "Washington DC"

	fmt.Println("Updated Map: ", capital)
}
