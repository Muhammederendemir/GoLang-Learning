package main

import "fmt"

func main() {

	// declare a struct Person
	type Person struct {
		name string
		age  int
	}

	// instance of the struct Person
	person1 := Person{"Muhammed Eren", 24}

	// create a struct type pointer that
	// stores the address of person1
	var ptr *Person
	ptr = &person1

	// print struct instance
	fmt.Println(person1)

	// print the struct type pointer
	fmt.Println(ptr)

}
