package main

import "fmt"

func main() {

	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")

}
