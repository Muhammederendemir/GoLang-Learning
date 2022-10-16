package main

import "fmt"

func division(num1, num2 int) {

	// if num2 is 0, program is terminated due to panic
	if num2 == 0 {
		panic("Cannot divide a number by zero")
	} else {
		result := num1 / num2
		fmt.Println("Result: ", result)
	}
}

func main() {

	division(4, 2)
	division(8, 0)
	division(2, 8)

}
