// Program using range with map

package main

import "fmt"

func main() {

	// create a map
	subjectMarks := map[string]float32{"Java": 80, "Python": 81, "Golang": 85}
	fmt.Println("Marks obtained:")

	// use for range to iterate through the key-value pair
	for subject, marks := range subjectMarks {
		fmt.Println(subject, ":", marks)
	}

}
