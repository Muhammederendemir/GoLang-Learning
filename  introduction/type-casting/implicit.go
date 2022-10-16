package main

import "fmt"

func main() {

	// initialize integer variable to a floating-point number

	//var number int = 4.34

	//error : ./prog.go:8:7: constant 4.34 truncated to integer

	var number int = 4

	fmt.Printf("Number is %d", number)
}
