// If package is called main, it will create an executable
package main

import "fmt"

// There can only be one main functions, so that Golang knows which main to call
func main() {
	learnSlices()
}

func learnVars() {
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

func learnFmt() {
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

func learnArrays() {
	var explicitArr [3]int = [3]int{1, 2, 3}
	implicitArr := [3]int{1, 2, 3}

	fmt.Println(explicitArr[1] + implicitArr[2])
}

func learnSlices() {
	var explicitSlice []int = []int{1, 2, 3}
	implicitSlice := []int{1, 2, 3}
	explicitSlice = append(explicitSlice, 4)
	implicitSlice = append([]int{0}, implicitSlice...)
	fmt.Printf("Explicit slice: %v\n", explicitSlice)
	fmt.Printf("Implicit slice: %v\n", implicitSlice)

	// slice ranges: inclusive of first number but not the second
	rangeSlice := implicitSlice[0:3]
	fmt.Printf("Range slice: %v\n", rangeSlice)
	startRange := implicitSlice[:3]
	fmt.Printf("Start range slice: %v\n", startRange)
	endRange := implicitSlice[1:]
	fmt.Printf("End range slice: %v\n", endRange)

}
