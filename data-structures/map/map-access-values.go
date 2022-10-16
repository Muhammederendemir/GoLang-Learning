package main

import "fmt"

func main() {

	// create a map
	flowerColor := map[string]string{"Sunflower": "Yellow", "Jasmine": "White", "Hibiscus": "Red"}

	// access value for key Sunflower
	fmt.Println(flowerColor["Sunflower"]) // Yellow

	// access value for key Hibiscus
	fmt.Println(flowerColor["Hibiscus"]) // Red

}
