// If package is called main, it will create an executable
package main

import "fmt"

// There can only be one main functions, so that Golang knows which main to call
func main() {
	learn_fmt()
}

func learn_vars() {
	// string: ------------------------------------------------
	// Declare variables: var {name} {type}
	var explicitName string = "explicit string"
	fmt.Println(explicitName)

	var implicitName = "implicit string"
	fmt.Println(implicitName)

	explicitName = "explicit string 2"

	// Do := only for initializing, after which use =
	// This cannot be used outside a function
	nameWithVar := "No var string"
	fmt.Println(nameWithVar)

	nameWithVar = "No var updated string"

	// int: ----------------------------------------------------
	var explicitInt int = 1
	var implicitInt = 2
	noVarInt := 3
	fmt.Println(explicitInt + implicitInt + noVarInt)

	// typesOfInt: (no. of bits used) --------------------------
	var num8Bit int8 = 127
	num8Bit += 1
	fmt.Println(num8Bit) // Overflowed

	var num8Unsigned uint8 = 255
	fmt.Println(num8Unsigned)
}

func learn_fmt() {
	name := "mong"
	age := "24"

	fmt.Println("String with commas", name, " age: ", age)

	// Formatted string --> Printf
	// %v is a format specifier for variables
	fmt.Printf("Formatted: My name: %v and my age: %v\n", name, age)
	// %q is to add quotes
	fmt.Printf("My name is %q\n", name)
	// %T is the type of variable
	fmt.Printf("%v is of type %T\n", name, name)
	// %f is for float
	fmt.Printf("Food cost $%0.2f\n", 0.50)

	// Save formatted string --> Sprintf
	savedStr := fmt.Sprintf("Saved: age: %v, name: %v", age, name)
	fmt.Println(savedStr)
}
