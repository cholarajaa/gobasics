// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare an array of 5 strings with each element initialized to its zero value.
//
// Declare a second array of 5 strings and initialize this array with literal string
// values. Assign the second array to the first and display the results of the first array.
// Display the string value and address of each element.
package main

import "fmt"

// Add imports.

func main() {
	// var array [5]int
	var array [5]string
	// Declare an array of 5 strings set to its zero value.
	fmt.Printf("%v \t %d\n", array, cap(array))
	// Declare an array of 5 strings and pre-populate it with names.
	// name_array := [5]string{"madhu", "sneha", "Amulya", "Kushal", "vittu"}
	name_array := [5]string{"vittu", "advika", "gouthami", "pratheeth", "Shanu"}
	// Assign the populated array to the array of zero values.
	fmt.Printf("%v \t %d\n", name_array, cap(name_array))

	array = name_array
	// Iterate over the first array declared.
	for _, name := range name_array {
		// Display the string value and address of each element.
		fmt.Printf("[Addr: %p, Value: %v]\n", &name, name)
	}
}
