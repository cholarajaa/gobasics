// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare a nil slice of integers. Create a loop that appends 10 values to the
// slice. Iterate over the slice and display each value.
//
// Declare a slice of five strings and initialize the slice with string literal
// values. Display all the elements. Take a slice of index one and two
// and display the index position and value of each element in the new slice.
package main

import "fmt"

// Add imports.

func main() {

	// Declare a nil slice of integers.
	var nil []int

	// Append numbers to the slice.
	nil = append(nil, 1)
	nil = append(nil, 2)
	nil = append(nil, 3)

	// Display each value in the slice.
	for _, value := range nil {
		fmt.Println(value)
	}

	// Declare a slice of strings and populate the slice with names.
	var names = []string{"vittu", "Pratheeth", "Shantu", "Rishika", "Advika"}

	// names = []string{"vittu", "Pratheeth", "Shantu"}

	// Display each index position and slice value.
	for index, name := range names {
		fmt.Printf("Index: %d\t Value: %v\n", index, name)
	}

	// Take a slice of index 1 and 2 of the slice of strings.
	names1 := names[1:2]

	// Display each index position and slice values for the new slice.
	for index, name := range names1 {
		fmt.Printf("Index: %d\t Value: %s\n", index, name)
	}
}
