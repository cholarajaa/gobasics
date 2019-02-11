// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Declare three variables that are initialized to their zero value and three
// declared with a literal value. Declare variables of type string, int and
// bool. Display the values of those variables.
//
// Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).
package main

import "fmt"

// Add imports

// main is the entry point for the application.
func main() {

	// Declare variables that are set to their zero value.
	var a int
	var b string
	var c bool
	var d []string
	// Display the value of those variables.
	fmt.Printf("Type %T\t value %v\n", a, a)
	fmt.Printf("Type %T\t value %v\n", b, b)
	fmt.Printf("Type %T\t value %v\n", c, c)
	fmt.Printf("Type %T\t value %v\n\n", d, d)

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	aa := 64
	bb := "sweet"
	cc := true
	dd := []string{"test", "this", "syntax"}

	// Display the value of those variables.
	fmt.Printf("Type %T\t value %v\n", aa, aa)
	fmt.Printf("Type %T\t value %v\n", bb, bb)
	fmt.Printf("Type %T\t value %v\n", cc, cc)
	fmt.Printf("Type %T\t value %v\n", dd, dd)
	// Perform a type conversion.
	// Display the value of that variable.
	ab := string(aa)
	fmt.Printf("Type %T\t value %v\n\n", ab, ab)
	fmt.Printf("%T\n", aa)
}
