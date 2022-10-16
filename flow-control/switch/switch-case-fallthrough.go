// Program to print the day of the week using fallthrough in switch

package main

import "fmt"

func main() {
	dayOfWeek := 3

	switch dayOfWeek {
	case 1:
		fmt.Println("Sunday")

	case 2:
		fmt.Println("Monday")

	case 3:
		fmt.Println("Tuesday")
		fallthrough

	case 4:
		fmt.Println("Wednesday")

	case 5:
		fmt.Println("Thursday")

	case 6:
		fmt.Println("Friday")

	case 7:
		fmt.Println("Saturday")

	default:
		fmt.Println("Invalid day")
	}
}
