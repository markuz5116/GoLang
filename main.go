// If package is called main, it will create an executable
package main

import "fmt"

// There can only be one main functions, so that Golang knows which main to call
func main() {
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
}
