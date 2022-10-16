package main

// import multiple packages
import (
	"fmt"
	"strings"
)

func main() {

	// convert the string to lowercase
	lower := strings.ToLower("GOLANG STRINGS")
	fmt.Println(lower)

	// convert the string to uppercase
	upper := strings.ToUpper("golang strings")
	fmt.Println(upper)

	// create a string array
	stringArray := []string{"I love", "Go Programming"}

	// join elements of array with space in between
	joinedString := strings.Join(stringArray, " ")
	fmt.Println(joinedString)

}
