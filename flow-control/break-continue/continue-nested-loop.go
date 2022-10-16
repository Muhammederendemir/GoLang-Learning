package main

import "fmt"

func main() {
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {

			// skips the inner for loop only
			if j == 2 {
				continue
			}

			fmt.Println("i=", i, "j=", j)

		}
	}
}
