package main

import "fmt"

func learnPointer() {
	originalStr := "mong"
	fmt.Println("Original:", originalStr)

	updateCopy(originalStr)
	fmt.Println("\"Updated\" name:", originalStr)

	updatePointerVal(&originalStr)
	fmt.Println("Actually updated name:", originalStr)
}

func updateCopy(str string) {
	str = "new string"
}

func updatePointerVal(str *string) {
	*str = "new string"
}
