// If package is called main, it will create an executable
package main

import "fmt"

// There can only be one main functions, so that Golang knows which main to call
func main() {
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
