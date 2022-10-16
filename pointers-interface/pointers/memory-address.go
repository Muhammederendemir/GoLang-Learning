// Program to illustrate how memory address works

package main

import "fmt"

func main() {

	var num int = 5
	// prints the value stored in variable
	fmt.Println("Variable Value:", num)

	// prints the address of the variable
	fmt.Println("Memory Address:", &num)

}
