package main

import "fmt"

// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

// Struct types are a way of creating
// complex types that group fields of data together.

// Declare a struct type to maintain information about a user (name, email and age).
// Create a value of this type, initialize with values and display each field.
//
// Declare and initialize an anonymous struct type with the same three fields. Display the value.

// Add imports.

// Add user type and provide comment.
type user struct {
	name  string
	email string
	age   int
}

func main() {

	// Declare variable of type user and init using a struct literal.
	u1 := user{
		name:  "madhu",
		email: "madhureddy88@gmail.com",
		age:   1,
	}
	// Display the field values.
	fmt.Printf("name: %s\n", u1.name)
	fmt.Printf("email: %s\n", u1.email)
	fmt.Printf("age: %d\n\n", u1.age)

	// Declare a variable using an anonymous struct.
	u2 := struct {
		name  string
		email string
		age   int
	}{
		name:  "madhu",
		email: "madhureddy88@gmail.com",
		age:   10,
	}

	// Display the field values.
	fmt.Printf("name: %s\n", u2.name)
	fmt.Printf("email: %s\n", u2.email)
	fmt.Printf("age: %d\n", u2.age)
}
