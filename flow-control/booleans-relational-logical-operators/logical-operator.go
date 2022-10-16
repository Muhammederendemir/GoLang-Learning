// Program to illustrate the working of Logical Operator

package main

import "fmt"

func main() {

	number1 := 6
	number2 := 12
	number3 := 6
	var result bool

	// returns false because number1 > number2 is false
	result = (number1 > number2) && (number1 == number3)

	fmt.Printf("Result of AND operator is %t \n", result)

	// returns true because number1 == number3 is true
	result = (number1 > number2) || (number1 == number3)

	fmt.Printf("Result of OR operator is %t \n", result)

	// returns false because number1 == number3 is true
	result = !(number1 == number3)

	fmt.Printf("Result of NOT operator is %t \n", result)

}
