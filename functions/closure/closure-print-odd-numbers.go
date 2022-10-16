package main

import "fmt"

func calculate() func() int {
	num := 1

	// returns inner function
	return func() int {
		num = num + 2
		return num
	}

}

func main() {

	// call the outer function
	odd := calculate()

	// call the inner function
	fmt.Println(odd())
	fmt.Println(odd())
	fmt.Println(odd())

	// call the outer function again
	odd2 := calculate()
	fmt.Println(odd2())

}
