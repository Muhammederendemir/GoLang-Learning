// Program to retrieve the keys of a map

package main

import "fmt"

func main() {

	// create a map
	subjectMarks := map[string]float32{"Java": 80, "Python": 81, "Golang": 85}

	fmt.Println("Subjects:")
	for subject := range subjectMarks {
		fmt.Println(subject)
	}
}
