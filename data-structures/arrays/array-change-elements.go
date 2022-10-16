// Program to change element by assigning the desired index with a new value

package main

import "fmt"

func main() {
	weather := [3]string{"Rainy", "Sunny", "Cloudy"}

	// change the element of index 2
	weather[2] = "Stromy"

	fmt.Println(weather)
}
