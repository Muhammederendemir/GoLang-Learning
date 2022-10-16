// Program to illustrate call by reference

package main

import "fmt"

// call by value
func callByValue(num int) {

	num = 30
	fmt.Println(num) // 30

}

// call by reference
func callByReference(num *int) {

	*num = 10
	fmt.Println(*num) // 10

}

func main() {

	var number int

	// passing value
	callByValue(number)

	// passing a reference (address)
	callByReference(&number)

}
