// Program to end the recursive function using if…else

package main

import "fmt"

func countDown(number int) {

	if number > 0 {
		fmt.Println(number)

		// recursive call
		countDown(number - 1)
	} else {
		// ends the recursive function
		fmt.Println("Countdown Stops")
	}

}

func main() {
	countDown(3)
}
