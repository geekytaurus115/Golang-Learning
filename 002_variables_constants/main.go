package main

import "fmt"

func main() {

	/*
			Variables:
			--> Mutable, can be reassigned.
		    --> Use the `var` keyword or
			--> short declaration := for implicit type inference.

		    Constants:
			--> Immutable, cannot be reassigned.
		    --> Use the `const` keyword.
	*/

	// ******** Variables *********

	// Declaration
	var myVar string

	// Initialization
	var myName string = "Geekytaurus"

	// Type Inference
	// --> Go can infer the variable type if it's assigned a value during declaration.
	country := "INDIA"

	// Zero Values
	// --> Variables without an explicit value are assigned a zero value based on
	// --> their type (0 for numeric types, "" for strings, nil for pointers, etc.)
	var score float64  // score is assigned a zero value of 0.0
	fmt.Println(score) //0.0

	var testString string
	fmt.Println(testString) // ""

	// Multiple Variable Declaration - Multiple variable can be declared in a single statement
	var x, y, z int

	// *************** CONSTANTS ****************

	// Declaration
	const PI = 3.1415

	// Typed Constants - constant can have explicit types
	const DAYS int = 7

	// Enumerated Constants - Enumerated constants are declared using the iota identifier for incremental values.
	const (
		Monday = iota // 0
		Tuesday
		Wednesday
	)

	// Typed Enumeration - Enumerated constants with typed values
	const (
		January int = iota + 1 //1
		February
		March
	)

	// Expression Constants
	const (
		X = 10
		Y = X + 5
	)

	/*
			```
		  	  p := 6
			  const Z = p + 2
		    ```
			ERROR: p + 2 (value of type int) is not constant
	*/

}
