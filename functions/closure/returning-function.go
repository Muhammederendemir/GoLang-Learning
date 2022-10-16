package main

import "fmt"

func greet() func() {

	return func() {
		fmt.Println("Hi Muhammed Eren")
	}

}

func main() {

	g1 := greet()
	g1()
}
