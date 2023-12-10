package main

import (
	"fmt"
)

// Function variadic (using slicing)
func sum(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

// Function variadic (using slicing)
func infoNames(values ...string) {
	for _, value := range values {
		fmt.Println(value)
	}
}

// Names return values
func getValues(value int) (double int, triple int, quad int) {
	double = value * 2
	triple = value * 3
	quad = value * 4
	return
}

func main() {

	// Function can return multiple values (Function variadic)
	// https://golang.org/ref/spec#Passing_arguments_to_..._parameters
	fmt.Println(sum(1, 2, 3, 4, 5))

	infoNames("John", "Doe")

	fmt.Println(getValues(2))

	// Generate a slice of numbers and strings
	slice := []interface{}{1, "John", 2, "Doe"}
	fmt.Println(slice)
}
