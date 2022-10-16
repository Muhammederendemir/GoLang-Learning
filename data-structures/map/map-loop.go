package main

import "fmt"

func main() {

	// create a map
	squaredNumber := map[int]int{2: 4, 3: 9, 4: 16, 5: 25}

	// for-range loop to iterate through each key-value of map
	for number, squared := range squaredNumber {
		fmt.Printf("Square of %d is %d\n", number, squared)
	}

}
