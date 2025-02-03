package main

import "fmt"

func main() {
	// Original array of 24 items
	s := 25
	originalArray := [25]int{}
	for i := 0; i < s; i++ {
		originalArray[i] = i + 1
	}

	// Loop to create arrays of max size 10 items
	for i := 0; i < len(originalArray); i += 10 {
		// end := i + 10
		// if end > len(originalArray) {
		// 	end = len(originalArray)
		// }
		subArray := originalArray[i:min(i+10, len(originalArray))]
		fmt.Println(subArray)
	}
}
