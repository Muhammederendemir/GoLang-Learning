// Program to illustrate the working of Relational Operators

package main

import "fmt"

func main() {

	number1 := 12
	number2 := 20
	var result bool

	// equal to operator
	result = (number1 == number2)

	fmt.Printf("%d == %d returns %t \n", number1, number2, result)

	// not equal to operator
	result = (number1 != number2)

	fmt.Printf("%d != %d returns %t \n", number1, number2, result)

	// greater than operator
	result = (number1 > number2)

	fmt.Printf("%d > %d returns %t \n", number1, number2, result)

	// less than operator
	result = (number1 < number2)

	fmt.Printf("%d < %d returns %t \n", number1, number2, result)

}
